package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro/v2"
    hello  "clientdemo/proto/clientdemo"
)

func main() {
    // New Service
    service := micro.NewService(
        micro.Name("service.hello-world.client"), //name the client service
    )
    // Initialise service
    service.Init()

    //create hello service client
    helloClient := hello.NewServerdemoService("cn.cmis110.service.serverdemo", service.Client())

    //invoke hello service method
    resp, err := helloClient.HelloWorld(context.TODO(), &hello.Request{Name: "xiaoming!"})
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf(" code: %+v \n message:%+v \n",resp.Code,resp.Message)
}