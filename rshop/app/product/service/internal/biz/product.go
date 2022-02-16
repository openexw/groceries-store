package biz

import (
	"context"
	"github.com/google/wire"
	"log"
)

type ProductRepo interface {
	ListProduct(ctx context.Context, pageNum, pageSize, channel int64) ([]*Product, error)
}

var ProviderSet = wire.NewSet(NewProductUserCase)

// biz.Product 用于业务中的显示，用于 BFF 层

type Product struct {
	ID          int64
	ProductCode string
	Name        string
	Channel     int64
	Period      int64
	WaresType   int64
	Unit        int64
	TotalAmount float64
	Desc        string
}

type ProductUserCase struct {
	repo ProductRepo
	log  *log.Logger
}

func NewProductUserCase(repo ProductRepo, log *log.Logger) *ProductUserCase {
	return &ProductUserCase{repo: repo, log: log}
}

func (puc *ProductUserCase) List(ctx context.Context, pageNum, pageSize, channel int64) ([]*Product, error) {
	return puc.repo.ListProduct(ctx, pageNum, pageSize, channel)
}
