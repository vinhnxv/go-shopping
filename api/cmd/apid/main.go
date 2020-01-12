package main

import (
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	_ "github.com/micro/go-micro/client/grpc"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"github.com/vinhnxv/go-shopping/api/internal/platform/config"
	"github.com/vinhnxv/go-shopping/api/internal/service"
)

const (
	serviceName = "go.shopping.api.v1.commerce"
)

func main() {
	webService := web.NewService(
		web.Name(serviceName),
		web.Version(config.Version),
	)

	if err := webService.Init(); err != nil {
		log.Fatal(err)
	}
	handler := service.NewCommerceService(client.DefaultClient)

	wc := restful.NewContainer()
	ws := new(restful.WebService)

	ws.
		Path("/v1/commerce").
		Doc("Aggregated Commerce API").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/products/{sku}").To(handler.GetProductDetails)).
		Doc("Query product details")

	wc.Add(ws)
	webService.Handle("/", wc)
	if err := webService.Run(); err != nil {
		log.Fatal(err)
	}
}
