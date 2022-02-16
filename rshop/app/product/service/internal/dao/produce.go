package dao

import (
	"context"
	"log"
	"rshop/app/product/service/internal/biz"
	"rshop/pkg"
	"time"
)

type Product struct {
	ID          int64     `gorm:"column:id" json:"id"`
	ProductCode string    `gorm:"column:product_code" json:"product_code"`
	Name        string    `gorm:"column:name" json:"name"`
	Channel     int64     `gorm:"column:channel" json:"channel"`
	Period      int64     `gorm:"column:period" json:"period"`
	WaresType   int64     `gorm:"column:wares_type" json:"wares_type"`
	Unit        int64     `gorm:"column:unit" json:"unit"`
	TotalAmount float64   `gorm:"column:total_amount" json:"total_amount"`
	Desc        string    `gorm:"column:desc" json:"desc"`
	Created     time.Time `gorm:"column:created" json:"created"`
	Updated     time.Time `gorm:"column:updated" json:"updated"`
}

type productRepo struct {
	dao *Dao
	log *log.Logger
}

func NewProductRepo(dao *Dao, log *log.Logger) biz.ProductRepo {
	return &productRepo{dao: dao, log: log}
}

func (pr *productRepo) ListProduct(ctx context.Context, pageNum, pageSize, channel int64) ([]*biz.Product, error) {
	var p []Product

	rst := pr.dao.db.WithContext(ctx).
		Limit(int(pageSize)).
		Where("channel=?", channel).
		Offset(int(pkg.GetPageOffset(pageNum, pageSize))).
		Find(&p)
	if rst.Error != nil {
		return nil, rst.Error
	}

	bizProducts := make([]*biz.Product, rst.RowsAffected)
	for _, item := range p {
		bizProducts = append(bizProducts, &biz.Product{
			ProductCode: item.ProductCode,
			Name:        item.Name,
			Channel:     item.Channel,
			Period:      item.Period,
			WaresType:   item.WaresType,
			Unit:        item.Unit,
			TotalAmount: item.TotalAmount,
			Desc:        item.Desc,
		})
	}

	return bizProducts, nil
}
