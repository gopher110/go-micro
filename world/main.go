package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro/v2"
    hello  "hello/proto/hello"
    "log"
    "time"
)

func main() {
    // New Service
    service := micro.NewService(
        micro.Name("service.hello-world.client"), //name the client service
    )
    // Initialise service
    service.Init()
    // create publisher
    pub := micro.NewPublisher("service.hello-world", service.Client())

    // publish message every second
    for now := range time.Tick(2*time.Second) {
        if err := pub.Publish(context.TODO(), &hello.Message{Say: fmt.Sprintf("**%s**",now.String())}); err != nil {
            log.Fatal("publish err", err)
        }
    }
    //create hello service client
    helloClient := hello.NewHelloService("service.hello-world", service.Client())

    //invoke hello service method
    resp, err := helloClient.Call(context.TODO(), &hello.Request{Name: " world!"})
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(resp.Msg)
}