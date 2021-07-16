package main

import (
	"log"
	"time"
	"math/rand"
	"encoding/json"
	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

func main(){
	// Instantiate a producer.
	config := nsq.NewConfig()
	var err error
	producer, err = nsq.NewProducer("10.32.5.119:31726", config)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().Unix())
	
	for j:=0; j<10000; j++{
		loopN := rand.Intn(5)
		for i := 0; i < loopN; i++{
			push()
		}
		log.Printf("sent times: %d", loopN)
		time.Sleep(time.Duration(100)*time.Second)
	}

	// Gracefully stop the producer.
	producer.Stop()
}


func push(){
	requestData := []request{
		{
			ApiUri:"https://hq.sinajs.cn/list=sz002023",
			Memthod:"GET",
		},
		{
			ApiUri:"https://hq.sinajs.cn/list=sh600009",
			Memthod:"GET",
			MsgData: `{"id":"123,456"}`,
		},
		{
			ApiUri:"https://hq.sinajs.cn/list=sh600096",
			Memthod:"GET",
			MsgData: `{"v":"v3"}`,
		},
		{
			ApiUri:"https://hq.sinajs.cn/list=sz002023",
			Memthod:"GET",
		},
		{
			ApiUri:"https://hq.sinajs.cn/list=sz300059",
			Memthod:"GET",
		},
		{
			ApiUri:"https://hq.sinajs.cn/list=sh601111",
			Memthod:"GET",
			MsgData: `{"id":"4125"}`,
		},
		{
			ApiUri:"https://ms.yimishiji.com/monitor-api/wxwork-message/create",
			Memthod:"POST",
			MsgData: `{"access_token":"4125","totag":"17","title":"nsq管道测试"}`,
		},
	}
	
	messageBody, err := json.Marshal(requestData[rand.Intn(len(requestData))])

	if rand.Intn(100) == 1 {
		messageBody= []byte("this error json data, to test nsq")
	}

	//messageBody := []byte("hello push")
	topicName := "deliver-producer-api"

	// Synchronously publish a single message to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	err = producer.Publish(topicName, messageBody)
	if err != nil {
		log.Fatal(err)
	}
}

type request struct{
	ApiUri string
	Memthod string
	MsgData string
}