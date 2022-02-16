package service

import (
	"context"
	"github.com/google/wire"
	"log"
	v1 "rshop/api/product/service/v1"
	"rshop/app/product/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewProductService)

// 实现 grpc.pd,go 的方法

type ProductService struct {
	v1.UnimplementedProductServer

	biz *biz.ProductUserCase

	log *log.Logger
}

func NewProductService(biz *biz.ProductUserCase, log *log.Logger) *ProductService {
	return &ProductService{biz: biz, log: log}
}

func (ps ProductService) ListProduct(ctx context.Context, req *v1.ListProductReq) (*v1.ListProductReply, error) {
	rv, err := ps.biz.List(ctx, int64(req.PageNum), int64(req.PageSize), int64(req.Channel))
	if err != nil {
		return nil, err
	}

	rs := make([]*v1.ListProductReply_Product, 0)
	for _, x := range rv {
		if x == nil {
			continue
		}
		rs = append(rs, &v1.ListProductReply_Product{
			ProductCode: x.ProductCode,
			Name:        x.Name,
			Channel:     int32(x.Channel),
			Period:      int32(x.Period),
			WaresType:   int32(x.WaresType),
			Unit:        int32(x.Unit),
			TotalAmount: float32(int32(x.TotalAmount)),
			Desc:        x.Desc,
		})
	}
	return &v1.ListProductReply{
		Products: rs,
	}, nil
}
