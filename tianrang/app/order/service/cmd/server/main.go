package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	v1 "tianrang/api/order/service/v1"
	"tianrang/app/order/service/internal/biz"
	"tianrang/app/order/service/internal/conf"
	"tianrang/app/order/service/internal/dao"
	"tianrang/app/order/service/internal/server"
	"tianrang/app/order/service/internal/service"
)

var (
	Name    = "tianrang.order.service"
	Version string

	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "/Users/7yue/Workspace/go/src/other/groceries-store/tianrang/app/order/service/configs/config.yaml", "config path: eg: -conf config.yaml")
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
	var conf conf.Config
	if err := c.Scan(&conf); err != nil {
		panic(err)
	}

	ingredientRepo := dao.NewIngredientRepo()
	productRepo := dao.NewProductRepo()

	productBiz := biz.NewProductBiz(productRepo)
	ingredientBiz := biz.NewIngredientBiz(ingredientRepo, conf.Rules)

	srv := service.NewOrderService(productBiz, ingredientBiz, &conf)
	//example(srv)

	go func() {
		startPProf()
	}()
	runRpcServer(conf, srv)
}

func runRpcServer(conf conf.Config, srv *service.OrderService) {
	rpcServer := server.NewGRPCServer(conf.Server, srv)
	listener, err := net.Listen("tcp", conf.Server.Grpc.Addr)
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	log.Println("run :", listener.Addr())
	go http.ListenAndServe(":8080", nil)
	log.Fatal(rpcServer.Serve(listener))
}

func example(srv *service.OrderService) {
	ctx := context.TODO()

	items := []*v1.AddOrderReq_Items{
		{
			ProductCode:     "coconut_milk_tea",
			IngredientCodes: []string{"red_bean"},
		}, {
			ProductCode:     "coconut_milk_tea",
			IngredientCodes: []string{"taro_balls"},
		}, {
			ProductCode:     "sago_milk_tea",
			IngredientCodes: []string{"red_bean"},
		},
		// 当 cafe 添加非奶盖的情况
		//{
		//		ProductCode:     "cafe_americano",
		//		IngredientCodes: []string{"red_bean"},
		//	},
		{
			ProductCode:     "cafe_americano",
			IngredientCodes: []string{"cheese_milk_cover"},
		},
	}
	req := &v1.AddOrderReq{Items: items}
	_, err := srv.AddOrder(ctx, req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%v", srv.GetDesc())
	fmt.Printf("\n总价为：￥%.2f", srv.GetPrice())
}

func startPProf() {
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
