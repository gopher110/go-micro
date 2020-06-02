package main

import (
    "github.com/micro/go-micro/v2"
    log "github.com/micro/go-micro/v2/logger"
    "github.com/micro/go-micro/v2/server"
    "serverdemo/handler"
    serverdemo "serverdemo/proto/serverdemo"
    "serverdemo/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("service.pubsub"),
		micro.Version("latest"),
		micro.Address(":60002"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	serverdemo.RegisterServerdemoHandler(service.Server(), new(handler.Serverdemo))

    micro.RegisterSubscriber("service.pubsub", service.Server(), subscriber.Handler, server.SubscriberQueue("service.pubsub"))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
