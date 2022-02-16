package server

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	v1 "rshop/api/product/service/v1"
	"rshop/app/product/service/internal/conf"
	"rshop/app/product/service/internal/service"
)

var ProviderSet = wire.NewSet(NewGRPCServer)

func NewGRPCServer(c *conf.Server, s *service.ProductService) *grpc.Server {
	rpcServer := grpc.NewServer(grpc.ConnectionTimeout(c.Grpc.Timeout.AsDuration()))
	//if c.Grpc.
	v1.RegisterProductServer(rpcServer, s)

	return rpcServer
}
