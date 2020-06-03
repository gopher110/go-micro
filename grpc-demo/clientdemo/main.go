package main

import (
	hello "clientdemo/proto/clientdemo"
	"context"
	"fmt"

	"net/http"
	"sync"
	"time"

	hystrixGo "github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("service.hello-world.client"),     //name the client service
		micro.WrapClient(hystrix.NewClientWrapper()), //客户端熔断 默认的超时时间是1000毫秒， 默认最大并发数是10 github.com/afex/hystrix-go/hystrix
	)
	hystrixGo.DefaultMaxConcurrent = 2
	// Initialise service
	service.Init(micro.AfterStart(func() error {
		done := make(chan interface{})
		defer close(done)
		cond := sync.NewCond(&sync.Mutex{})
		var wg sync.WaitGroup
		count := 15
		wg.Add(count)
		for i := 0; i < count; i++ {
			goFunc(cond, func() {
				callHelloWorld(service)
				wg.Done()
			})
		}
		cond.Broadcast()
		wg.Wait()
		return nil
	}))

	service.Run()

}

func callHelloWorld(srv micro.Service) {
	//create hello service client
	helloClient := hello.NewServerdemoService("cn.cmis110.service.serverdemo", srv.Client())

	//invoke hello service method
	resp, err := helloClient.HelloWorld(context.TODO(), &hello.Request{Name: time.Now().Format("15:04:05")})
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

func goFunc(c *sync.Cond, fn func()) {
	var goroutineRunning sync.WaitGroup
	goroutineRunning.Add(1)
	go func() {
		goroutineRunning.Done()
		c.L.Lock()
		defer c.L.Unlock()
		c.Wait()
		fn()
	}()
	goroutineRunning.Wait()
}
