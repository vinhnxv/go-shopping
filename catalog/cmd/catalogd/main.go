package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	log "github.com/sirupsen/logrus"
	"github.com/vinhnxv/go-shopping/catalog/internal/platform/config"
	"github.com/vinhnxv/go-shopping/catalog/internal/platform/redis"
	"github.com/vinhnxv/go-shopping/catalog/internal/service"
	"github.com/vinhnxv/go-shopping/catalog/proto"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	svc := grpc.NewService(
		micro.Name(config.ServiceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version(config.Version),
	)
	svc.Init()

	redisCatalogRepository := redis.NewRedisRepository(":6379")
	if err := catalog.RegisterCatalogHandler(svc.Server(), service.NewCatalogService(redisCatalogRepository)); err != nil {
		panic(err)
	}

	if err := svc.Run(); err != nil {
		panic(err)
	}
}
