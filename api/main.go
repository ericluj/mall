package main

import (
	"mall/api/handler"
	"mall/api/srv"
	"mall/proto/user"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.api.api"),
	)

	if err := service.Init(); err != nil {
		panic(err)
	}

	service.Options().Service.Init(micro.Transport(grpc.NewTransport()))
	userSrv := user.NewUserService("go.micro.srv.user", service.Options().Service.Client())

	srv.Init(userSrv)
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

	return router
}
