package server

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	v1 "tianrang/api/order/service/v1"
	"tianrang/app/order/service/internal/conf"
	"tianrang/app/order/service/internal/service"
)

var ProviderSet = wire.NewSet(NewGRPCServer)

func NewGRPCServer(c *conf.Server, s *service.OrderService) *grpc.Server {
	rpcServer := grpc.NewServer(grpc.ConnectionTimeout(c.Grpc.Timeout.AsDuration()))
	//if c.Grpc.
	v1.RegisterOrderServer(rpcServer, s)

	return rpcServer
}
