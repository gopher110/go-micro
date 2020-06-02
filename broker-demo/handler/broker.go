package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	broker "broker-demo/proto/broker"
)

type Broker struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Broker) Call(ctx context.Context, req *broker.Request, rsp *broker.Response) error {
	log.Info("Received Broker.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Broker) Stream(ctx context.Context, req *broker.StreamingRequest, stream broker.Broker_StreamStream) error {
	log.Infof("Received Broker.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&broker.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Broker) PingPong(ctx context.Context, stream broker.Broker_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&broker.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
