package main

import (
	"github.com/micro/go-micro"
	gmbroker "github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/service/grpc"
	"github.com/vinhnxv/go-shopping/shipping/internal/platform/broker"
	"github.com/vinhnxv/go-shopping/shipping/internal/platform/config"
	"github.com/vinhnxv/go-shopping/shipping/internal/platform/redis"
	"github.com/vinhnxv/go-shopping/shipping/internal/service"
	"github.com/vinhnxv/go-shopping/shipping/proto"
	shipping "github.com/vinhnxv/go-shopping/shipping/proto"
	"log"
	"time"
)

func main() {
	if err := gmbroker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := gmbroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	repo := redis.NewRedisRepository(":6379")
	publisher := broker.NewEventPublisher()
	svc := grpc.NewService(
		micro.Name(config.ServiceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version(config.Version),
	)
	svc.Init()

	if err := shipping.RegisterShippingHandler(svc.Server(), service.NewShippingService(repo, publisher)); err != nil {
		panic(err)
	}

	if err := svc.Run(); err != nil {
		panic(err)
	}
}
