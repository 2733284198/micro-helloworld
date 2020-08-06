package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/xiuxiubiu/micro-helloworld/handler"
	"github.com/xiuxiubiu/micro-helloworld/subscriber"

	microhelloworld "github.com/xiuxiubiu/micro-helloworld/proto/micro-helloworld"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.micro-helloworld"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	microhelloworld.RegisterMicroHelloworldHandler(service.Server(), new(handler.MicroHelloworld))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.micro-helloworld", service.Server(), new(subscriber.MicroHelloworld))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
