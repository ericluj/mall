package main

import (
	"mall/product/conf"
	"mall/product/dao"
	"mall/product/handler"
	"mall/product/model"
	"mall/proto/product"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/broker/nsq/v2"
	"github.com/micro/go-plugins/registry/etcd/v2"
)

var config conf.Config

func main() {
	log.Name("go.micro.srv.product")
	b := nsq.NewBroker(nsq.WithLookupdAddrs([]string{"127.0.0.1:4161"}))
	service := micro.NewService(
		micro.Name("go.micro.srv.product"),
		micro.Transport(grpc.NewTransport()),
		micro.Broker(b),
		micro.Registry(etcd.NewRegistry()),
		micro.Flags(
			&cli.StringFlag{
				Name:  "database_driver",
				Usage: "database driver",
				Value: "mysql",
			},
			&cli.StringFlag{
				Name:  "database_host",
				Usage: "database host",
				Value: "127.0.0.1",
			},
			&cli.IntFlag{
				Name:  "database_port",
				Usage: "database port",
				Value: 3306,
			},
			&cli.StringFlag{
				Name:  "database_user",
				Usage: "database user",
				Value: "root",
			},
			&cli.StringFlag{
				Name:  "database_password",
				Usage: "database password",
				Value: "my123456",
			},
			&cli.StringFlag{
				Name:  "database_dbname",
				Usage: "database dbname",
				Value: "mall_product",
			},
		),
	)
	service.Init(
		micro.Action(func(c *cli.Context) error {
			config.DB.DriverName = c.String("database_driver")
			config.DB.Host = c.String("database_host")
			config.DB.Port = c.Int("database_port")
			config.DB.User = c.String("database_user")
			config.DB.PassWord = c.String("database_password")
			config.DB.DBName = c.String("database_dbname")
			log.Infof("database_driver:%s,database_host:%s,database_port:%d,database_user:%s,database_password:%s,database_dbname:%s\n", config.DB.DriverName, config.DB.Host, config.DB.Port, config.DB.User, config.DB.PassWord, config.DB.DBName)
			return nil
		}),
		micro.BeforeStart(func() (err error) {
			if err = dao.Init(&config); err != nil {
				return
			}
			if err = dao.GetDao().Ping(); err != nil {
				return
			}
			return
		}),
		micro.BeforeStop(func() error {
			return dao.GetDao().Disconnect()
		}),
	)

	brok := service.Server().Options().Broker
	if err := brok.Connect(); err != nil {
		panic(err)
	}
	// 订阅消息
	_, err := brok.Subscribe(model.TopicOrder, handler.Order, func(opts *broker.SubscribeOptions) { opts.Queue = model.ChannelProduct })
	if err != nil {
		log.Error(err)
	}

	if err := product.RegisterProductHandler(service.Server(), new(handler.Product)); err != nil {
		panic(err)
	}

	if err := service.Run(); err != nil {
		panic(err)
	}
}
