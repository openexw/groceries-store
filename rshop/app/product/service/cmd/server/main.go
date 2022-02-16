package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"log"
	"net"
	"net/http"
	"rshop/app/product/service/internal/conf"

	_ "net/http/pprof"
)

var (
	Name    = "rshop.product.service"
	Version string

	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "", "config path: eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	// load config
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	logger := log.Default()

	// 初始化
	rpcServer := InitApp(logger, bc.Data, bc.Server)
	// run
	listener, err := net.Listen("tcp", bc.Server.Grpc.Addr)
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	log.Println("run :", listener.Addr())
	go http.ListenAndServe(":8080", nil)
	log.Fatal(rpcServer.Serve(listener))

}
