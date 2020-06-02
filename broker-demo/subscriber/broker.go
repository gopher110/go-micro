package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	broker "broker-demo/proto/broker"
)

type Broker struct{}

func (e *Broker) Handle(ctx context.Context, msg *broker.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *broker.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
