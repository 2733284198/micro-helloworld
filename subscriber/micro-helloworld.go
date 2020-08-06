package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	microhelloworld "github.com/xiuxiubiu/micro-helloworld/proto/micro-helloworld"
)

type MicroHelloworld struct{}

func (e *MicroHelloworld) Handle(ctx context.Context, msg *microhelloworld.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *microhelloworld.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
