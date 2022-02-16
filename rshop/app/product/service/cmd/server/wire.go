package main

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"log"
	"rshop/app/product/service/internal/biz"
	"rshop/app/product/service/internal/conf"
	"rshop/app/product/service/internal/dao"
	"rshop/app/product/service/internal/server"
	"rshop/app/product/service/internal/service"
)

//initApp init kratos application.
func initApp(*log.Logger, *conf.Data, *conf.Server) *grpc.Server {
	panic(wire.Build(dao.ProviderSet, biz.ProviderSet, server.ProviderSet, service.ProviderSet))
}
