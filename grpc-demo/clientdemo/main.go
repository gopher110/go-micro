package main

import (
	hello "clientdemo/proto/clientdemo"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	"net/http"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("service.hello-world.client"), //name the client service
	)
	// Initialise service
	service.Init()

	callHelloWorld(service)

	service.Run()
}

func callHelloWorld(srv micro.Service) {
	//create hello service client
	helloClient := hello.NewServerdemoService("cn.cmis110.service.serverdemo", srv.Client())

	//invoke hello service method
	resp, err := helloClient.HelloWorld(context.TODO(), &hello.Request{Name: ""})
	if err != nil {
		e := errors.Parse(err.Error())
		if e.Code == http.StatusBadRequest {
			// deal with bad request
		}
		fmt.Println(e)
		return
	}
	fmt.Printf(" code: %+v \n message:%+v \n", resp.Code, resp.Message)
}
