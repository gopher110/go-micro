package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/server"
    "hello/handler"
	"hello/subscriber"

	hello "hello/proto/hello"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("service.hello-world"),
		micro.Version("v1.0"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("service.hello-world", service.Server(), subscriber.Handler, server.SubscriberQueue("service.hello-world"))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
