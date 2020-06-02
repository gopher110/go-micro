package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	serverdemo "serverdemo/proto/serverdemo"
)

type Serverdemo struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Serverdemo) Call(ctx context.Context, req *serverdemo.Request, rsp *serverdemo.Response) error {
	log.Info("Received Serverdemo.Call request")
	rsp.Message="Hello " + req.Name

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Serverdemo) Stream(ctx context.Context, req *serverdemo.StreamingRequest, stream serverdemo.Serverdemo_StreamStream) error {
	log.Infof("Received Serverdemo.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&serverdemo.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Serverdemo) PingPong(ctx context.Context, stream serverdemo.Serverdemo_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&serverdemo.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
// Call is a single request handler called via client.Call or the generated client code
func (e *Serverdemo) HelloWorld(ctx context.Context, req *serverdemo.Request, rsp *serverdemo.Response) error {
    log.Info("Received Serverdemo.Call request")
    rsp.Message="Hello " + req.Name
    rsp.Code=1

    return nil
}
