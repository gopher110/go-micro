package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	serverdemo "serverdemo/proto/serverdemo"
)

type Serverdemo struct{}

func (e *Serverdemo) Handle(ctx context.Context, msg *serverdemo.Message) error {
	log.Info("**接收到消息**: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *serverdemo.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
