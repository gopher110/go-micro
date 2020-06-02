package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
)

var (
	topic = "cn.cmis110.topic"
)

type Article struct {
	Id         int64     `json:"id"`
	Title      string    `jon:"title"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func generateBrokerMessage(id int64) (retVal *broker.Message) {
	var msg broker.Message
	ta := new(Article)
	ta.Id, ta.Title, ta.Content = id, time.Now().Format("15:04:05"), time.Now().String()
	msg.Body, _ = json.Marshal(*ta)
	msg.Header = map[string]string{"id": fmt.Sprintf("%d", id)}
	retVal = &msg
	return
}

func pub(brk broker.Broker) {
	var i int64
	for range time.Tick(2 * time.Second) {
		// build a message
		msg := generateBrokerMessage(i)
		// publish it
		if err := brk.Publish(topic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] published message:", string(msg.Body))
		}
		i++
	}
}

func sub(brk broker.Broker) {
	// subscribe a topic with queue specified
	if _, err := brk.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	}, broker.Queue(topic)); err != nil { //broker.DisableAutoAck(), rabbitmq.DurableQueue()，)
		fmt.Println(err)
	}
}

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("broker.example"),
	)

	// Initialise service
	service.Init(micro.AfterStart(func() error {
		brk := service.Options().Broker
		if err := brk.Connect(); err != nil {
			log.Fatalf("Broker Connect error: %v", err)
		}
		go sub(brk)
		go pub(brk)
		return nil
	}))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
