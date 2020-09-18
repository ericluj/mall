package main

import (
	"log"

	"mall/api/handler"
	"mall/api/srv"
	"mall/lib"
	"mall/lib/tracer"
	"mall/proto/order"
	"mall/proto/product"
	"mall/proto/user"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/etcd/v2"
	wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

func main() {
	service := web.NewService(
		web.Name(lib.ServiceApiName),
		web.Address(":9090"),
	)

	//链路追踪
	t, io, err := tracer.NewTracer(lib.ServiceApiName, "127.0.0.1:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	if err := service.Init(); err != nil {
		panic(err)
	}

	service.Options().Service.Init(
		micro.Transport(grpc.NewTransport()),
		micro.Registry(etcd.NewRegistry()),
		micro.WrapHandler(wrapperTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	userSrv := user.NewUserService(lib.ServiceUserName, service.Options().Service.Client())
	productSrv := product.NewProductService(lib.ServiceProductName, service.Options().Service.Client())
	orderSrv := order.NewOrderService(lib.ServiceOrderName, service.Options().Service.Client())
	srv.Init(userSrv, productSrv, orderSrv)
	router := Router()
	service.Handle("/", router)

	if err := service.Run(); err != nil {
		panic(err)
	}
}

func Router() *gin.Engine {
	router := gin.Default()
	r := router.Group("/api/v1/user")
	{
		r.GET("/query", handler.QueryUser)
		r.POST("/create", handler.CreateUser)
	}

	r1 := router.Group("/api/v1/product")
	{
		r1.GET("/query", handler.QueryProduct)
		r1.POST("/create", handler.CreateProduct)
		r1.POST("/update", handler.UpdateProduct)
	}

	r2 := router.Group("/api/v1/order")
	{
		r2.POST("/create", handler.CreateOrder)
	}

	return router
}
