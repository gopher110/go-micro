package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro/v2"
    hello  "clientdemo/proto/clientdemo"
    "log"
    "time"
)

func main() {
    // New Service
    service := micro.NewService(
        micro.Name("service.pubsub.pub"), //name the client service
    )
    // Initialise service
    service.Init()

    //create hello service client
    helloClient := hello.NewServerdemoService("service.pubsub", service.Client())

    //invoke hello service method
    resp, err := helloClient.HelloWorld(context.TODO(), &hello.Request{Name: "xiaofang!"})
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf(" code: %+v \n message:%+v \n",resp.Code,resp.Message)

    pub := micro.NewPublisher("service.pubsub", service.Client())
    // publish message every second
    for now := range time.Tick(2*time.Second) {
        if err := pub.Publish(context.TODO(), &hello.Message{Say:"now:"+ now.Format("15:04:05")}); err != nil {
            log.Fatal("publish err", err)
        }
    }


}