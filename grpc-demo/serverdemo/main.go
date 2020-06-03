package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	limiter "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	"serverdemo/handler"
	serverdemo "serverdemo/proto/serverdemo"
)

func main() {
	// New Service

	service := micro.NewService(
		micro.Name("cn.cmis110.service.serverdemo"),
		micro.Version("latest"),
		micro.Address(":60001"),
		micro.WrapHandler(limiter.NewHandlerWrapper(50)), //服务器限流 QPS设定上限50
	)

	// Initialise service
	service.Init()

	// Register Handler
	serverdemo.RegisterServerdemoHandler(service.Server(), new(handler.Serverdemo))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
