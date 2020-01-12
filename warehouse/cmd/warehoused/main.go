package main

import (
	"github.com/micro/go-micro"
	gmbroker "github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
	"github.com/vinhnxv/go-shopping/shipping/proto"
	"github.com/vinhnxv/go-shopping/warehouse/internal/platform/broker"
	"github.com/vinhnxv/go-shopping/warehouse/internal/platform/config"
	"github.com/vinhnxv/go-shopping/warehouse/internal/platform/redis"
	"github.com/vinhnxv/go-shopping/warehouse/internal/service"
	"github.com/vinhnxv/go-shopping/warehouse/proto"
	"time"
)

func main() {

	if err := gmbroker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := gmbroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	itemShippedChannel := make(chan *shipping.ItemShippedEvent)
	if err := broker.CreateEventConsumer(itemShippedChannel); err != nil {
		log.Fatalf("Broker error: %v", err)
	}

	repo := redis.NewWarehouseRepository(":6379")

	svc := grpc.NewService(
		micro.Name(config.ServiceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version(config.Version),
	)
	svc.Init()

	if err := warehouse.RegisterWarehouseHandler(svc.Server(), service.NewWarehouseService(repo, itemShippedChannel)); err != nil {
		panic(err)
	}

	if err := svc.Run(); err != nil {
		panic(err)
	}
}
