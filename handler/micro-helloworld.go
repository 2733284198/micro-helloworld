package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	microhelloworld "github.com/xiuxiubiu/micro-helloworld/proto/micro-helloworld"
)

type MicroHelloworld struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *MicroHelloworld) Call(ctx context.Context, req *microhelloworld.Request, rsp *microhelloworld.Response) error {
	log.Info("Received MicroHelloworld.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *MicroHelloworld) Stream(ctx context.Context, req *microhelloworld.StreamingRequest, stream microhelloworld.MicroHelloworld_StreamStream) error {
	log.Infof("Received MicroHelloworld.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&microhelloworld.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *MicroHelloworld) PingPong(ctx context.Context, stream microhelloworld.MicroHelloworld_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&microhelloworld.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
