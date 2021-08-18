package main

import (
	"compress/gzip"
	"io"
	"encoding/json"
	"time"
	"fmt"
	"strings"
	"math/rand"
	//"errors"
	//"bytes"
	"net/url"
	"github.com/nsqio/go-nsq"
	"github.com/nsqio/nsq/internal/lg"
	"github.com/nsqio/nsq/internal/clusterinfo"
)

type RequestWorker struct {
	logf     	lg.AppLogFunc
	opts     	*Options
	restChannel *clusterinfo.RestChannel
	consumer 	*nsq.Consumer
	ci          *clusterinfo.ClusterInfo
	writer      io.Writer
	gzipWriter  *gzip.Writer
	logChan     chan *nsq.Message
	termChan 	chan bool
	hupChan  	chan bool
	// for rotation
	filename 	string
	openTime 	time.Time
	filesize 	int64
	rev      	uint
}

func NewRequestWorker(logf lg.AppLogFunc, opts *Options, restChannel *clusterinfo.RestChannel, cfg *nsq.Config, ci *clusterinfo.ClusterInfo) (*RequestWorker, error) {
	consumer, err := nsq.NewConsumer(restChannel.Topic, restChannel.Channel, cfg)
	if err != nil {
		return nil, err
	}

	logf(lg.INFO, "create new request workder: [%s/%s]", restChannel.Topic, restChannel.Channel)

	f := &RequestWorker{
		logf:           logf,
		opts:           opts,
		restChannel:    restChannel,
		consumer:       consumer,
		ci:             ci,
		logChan:        make(chan *nsq.Message, 1),
		termChan:       make(chan bool),
		hupChan:        make(chan bool),
	}
	consumer.AddHandler(f)

	err = consumer.ConnectToNSQDs(opts.NSQDTCPAddrs)
	if err != nil {
		return nil, err
	}

	err = consumer.ConnectToNSQLookupds(opts.NSQLookupdHTTPAddrs)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *RequestWorker) router() {
	exit := false
	for {
		select {
		case <-f.consumer.StopChan:
			//消费通道关闭
			exit = true
		case <-f.termChan:
			//程序中止
			f.consumer.Stop()
		}

		if exit {
			break
		}
	}
}



func (f *RequestWorker) HandleMessage(m *nsq.Message) error {
	m.DisableAutoResponse()

	var req request
	var requestErr error

	//time.Sleep(time.Duration(rand.Intn(2))*time.Second)
	req.Channel = f.restChannel.Channel
	req.Topic = f.restChannel.Topic
	req.ApiUri = f.restChannel.RestUrl
	req.ContentType = f.restChannel.ContentType
	req.Memthod = f.restChannel.Method
	req.MsgData = string(m.Body)

	f.logf(lg.INFO, "[%s/%s] request goto [%s]%s", f.restChannel.Topic, f.restChannel.Channel, f.restChannel.Method, f.restChannel.RestUrl)
	
	if f.restChannel.Method=="GET"{
		req.ResponseCode, req.ResponseData, requestErr = f.PublishGet(f.restChannel.RestUrl, m.Body)
	} else if f.restChannel.Method=="POST"{
		if contentType=="application/x-www-form-urlencoded" {
			paramsData := map[string]string{}
			if err := json.Unmarshal([]byte(req.MsgData),&paramsData); err != nil{
				req.JsonErr = err.Error()
				req.JsonErrBody = string(req.MsgData)
			}else{
				paramValue := url.Values{}
				for key,value := range paramsData{
					paramValue.Add(key,value)
				}
				req.ResponseCode, req.ResponseData, requestErr = f.PublishPostForm(f.restChannel.RestUrl, &paramValue)
			}
		}else{
			req.ResponseCode, req.ResponseData, requestErr = f.PublishPost(f.restChannel.RestUrl, f.restChannel.ContentType, &paramValue)

		}
	}

	if requestErr != nil{
		req.ResponseErr = requestErr.Error()
	}
	
	req.MessageID = fmt.Sprintf("%s", m.ID)
	req.Attempts = m.Attempts
	
	f.PushLog(&req)

	if requestErr != nil || (req.JsonErr == "ok" && req.ResponseCode != 200){
		f.logf(lg.ERROR, "[%s/%s] request, messageid: %s errors: %s, req.ApiUri: %s", f.restChannel.Topic, f.restChannel.Channel, m.ID, req.ResponseErr, req.ApiUri)
		m.Requeue(-1)
		return requestErr
	}

	m.Finish()
	return nil
}


type request struct{
	MessageID string
	Attempts  uint16
	Topic string
	Channel string
	ContentType string
	ApiUri string
	Memthod string
	MsgData string
	ResponseErr string
	ResponseCode int
	ResponseData string
	JsonErr string
	JsonErrBody string
}

func (p *RequestWorker) PublishPostForm(addr string,  params *url.Values) (code int, response string, err error) {
	buf := strings.NewReader(params.Encode())	
	contentType := "application/x-www-form-urlencoded"
	resp, err2 := HTTPPost(addr, contentType, buf)
	if err2 != nil {
		return code, response, err2
	}
	//io.Copy(ioutil.Discard, resp.Body)
	var respByte []byte
	respByte, err = io.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		p.logf(lg.ERROR, "got status code %d", resp.StatusCode)
	}
	return resp.StatusCode, string(respByte), err
}

func (p *RequestWorker) PublishPost(addr string, contentType string, msg []byte) (code int, response string, err error) {
	buf := bytes.NewBuffer(msg)

	resp, err2 := HTTPPost(addr, contentType, buf)
	if err2 != nil {
		return code, response, err2
	}
	//io.Copy(ioutil.Discard, resp.Body)
	var respByte []byte
	respByte, err = io.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		p.logf(lg.ERROR, "got status code %d", resp.StatusCode)
	}
	return resp.StatusCode, string(respByte), err
}

func (p *RequestWorker) PublishGet(addr string, msgBody []byte) (code int, response string, err error)  {
	endpoint := addr
	var msg string
	if json.Valid(msgBody){
        paramsData := map[string]string{}
        
        json.Unmarshal(msgBody,&paramsData)
        paramValue := url.Values{}

        for key,value := range paramsData{
            paramValue.Add(key,value)
        }
        msg = paramValue.Encode()
    }else{
        msg = url.QueryEscape(string(msgBody))
    }

	if strings.Contains(addr, "?"){
		endpoint = fmt.Sprintf("%s&%s", addr, msg)
	}else{
		endpoint = fmt.Sprintf("%s?%s", addr, msg)
	}
	
	resp, err2 := HTTPGet(endpoint)
	if err2 != nil {
		p.logf(lg.ERROR, "HTTPGet %s error %d",endpoint, err2)
		return code, response, err2
	}
	//io.Copy(ioutil.Discard, resp.Body)
	var respByte []byte
	respByte, err = io.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode != 200 {
		p.logf(lg.ERROR, "got status code %d", resp.StatusCode)
	}
	return resp.StatusCode, string(respByte), err
}


//记录日志，失败后最多重试10次
func (t *RequestWorker) PushLog(req *request) error{
	//
	producers,err := t.ci.GetLookupdProducers(t.opts.NSQLookupdHTTPAddrs)
	if err != nil {
		t.logf(lg.ERROR, "could not get producers: %s", err)
		return err
	}

	messageBody, err := json.Marshal(req)
	if err != nil{
		t.logf(lg.ERROR, "request struct json encode error: %s", err)
		return err
	}

	topicName := "logs-request"

	for i := 1; i < 10; i++ {
		rand.Seed(time.Now().Unix())
		config := nsq.NewConfig()
		producer, err := nsq.NewProducer(producers[rand.Intn(len(producers))].TCPAddress(),config)
		if err != nil {
			t.logf(lg.ERROR, "couldn't create log producer error: %s", err)
			continue
		}

		// Synchronously publish a single message to the specified topic.
		// Messages can also be sent asynchronously and/or in batches.
		err = producer.Publish(topicName, messageBody)
		if err != nil {
			t.logf(lg.ERROR, "[%s/%s] request message pulish error, errors: %s, messageBody: %s", t.restChannel.Topic, t.restChannel.Channel, err, messageBody)
			continue
		}else{
			break
		}
    } 
	return nil
}