package main

import (
	"compress/gzip"
	"io"
	"encoding/json"
	"time"
	"fmt"
	"strings"
	//"math/rand"
	//"errors"
	//"bytes"
	"net/url"
	"github.com/nsqio/go-nsq"
	"github.com/nsqio/nsq/internal/lg"
)

type RequestWorker struct {
	logf     lg.AppLogFunc
	opts     *Options
	topic    string
	consumer *nsq.Consumer
	producer *nsq.Producer

	writer         io.Writer
	gzipWriter     *gzip.Writer
	logChan        chan *nsq.Message

	termChan chan bool
	hupChan  chan bool

	// for rotation
	filename string
	openTime time.Time
	filesize int64
	rev      uint
}

func NewRequestWorker(logf lg.AppLogFunc, opts *Options, topic string, cfg *nsq.Config, producer *nsq.Producer) (*RequestWorker, error) {
	consumer, err := nsq.NewConsumer(topic, opts.Channel, cfg)
	if err != nil {
		return nil, err
	}

	f := &RequestWorker{
		logf:           logf,
		opts:           opts,
		topic:          topic,
		consumer:       consumer,
		producer:       producer,
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

	if err := json.Unmarshal(m.Body,&req); err != nil{
		req.JsonErr = err.Error()
		req.JsonErrBody = string(m.Body)
	}else{
		f.logf(lg.INFO, "[%s/%s] request goto [%s]%s", f.topic, f.opts.Channel, req.Memthod, req.ApiUri)
		
		paramsData := map[string]string{}
		if err = json.Unmarshal([]byte(req.MsgData),&paramsData); err != nil{
			req.JsonErr = err.Error()
			req.JsonErrBody = string(req.MsgData)
		}else{
			paramValue := url.Values{}
			for key,value := range paramsData{
				paramValue.Add(key,value)
			}

			if req.Memthod=="GET"{
				req.ResponseCode, req.ResponseData, requestErr = f.PublishGet(req.ApiUri, &paramValue)
			} else if req.Memthod=="POST"{
				req.ResponseCode, req.ResponseData, requestErr = f.PublishPost(req.ApiUri, &paramValue)
			}
		
		}

		if requestErr != nil{
			req.ResponseErr = requestErr.Error()
		}
		// if rand.Intn(50) == 1{
		// 	return errors.New("request time out")
		// }
	}
	
	req.Channel = f.opts.Channel
	req.Topic = f.topic
	req.MessageID = fmt.Sprintf("%s", m.ID)
	req.Attempts = m.Attempts
	messageBody, err := json.Marshal(req)
	//messageBody := []byte("hello push")
	topicName := "logs-request"
	// Synchronously publish a single message to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	err = f.producer.Publish(topicName, messageBody)
	if err != nil {
		f.logf(lg.ERROR, "[%s/%s] request message pulish error, errors: %s, messageBody: %s", f.topic, f.opts.Channel, err, messageBody)
		m.Requeue(-1)
		return err
	}

	if requestErr != nil || (req.JsonErr == "ok" && req.ResponseCode != 200){
		f.logf(lg.ERROR, "[%s/%s] request, messageid: %s errors: %s, req.ApiUri: %s", f.topic, f.opts.Channel, m.ID, req.ResponseErr, req.ApiUri)
		m.Requeue(-1)
		return err
	}

	m.Finish()
	return nil
}


type request struct{
	MessageID string
	Attempts  uint16
	Topic string
	Channel string
	ApiUri string
	Memthod string
	MsgData string
	ResponseErr string
	ResponseCode int
	ResponseData string
	JsonErr string
	JsonErrBody string
}


func (p *RequestWorker) PublishPost(addr string, params *url.Values) (code int, response string, err error) {
	buf := strings.NewReader(params.Encode())	
	resp, err2 := HTTPPost(addr, buf)
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

func (p *RequestWorker) PublishGet(addr string, params *url.Values) (code int, response string, err error)  {
	endpoint := addr
	msg := params.Encode()

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
