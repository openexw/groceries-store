package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/wire"
	v1 "tianrang/api/order/service/v1"
	"tianrang/app/order/service/internal/biz"
	"tianrang/app/order/service/internal/conf"
)

var ProviderSet = wire.NewSet(NewOrderService, NewOrderItem)
var (
	ErrIngredientsNotFound = errors.New("配料不存在")
	//ErrAddIngredientsRules = errors.New("添加配料错误")

	ErrProductNotFound = errors.New("产品不存在")

	ErrOrderOverflow = errors.New("订单超过最大限制")
)

// 实现 grpc.pd.go 的方法

type OrderService struct {
	v1.UnimplementedOrderServer
	productBiz    *biz.ProductBiz
	ingredientBiz *biz.IngredientBiz

	//log *log.Logger
	Order
}

func NewOrderService(productBiz *biz.ProductBiz, ingredientBiz *biz.IngredientBiz, conf *conf.Config) *OrderService {
	productBiz.LoadData([]byte(conf.Data.Products))
	ingredientBiz.LoadData([]byte(conf.Data.Ingredients))

	return &OrderService{productBiz: productBiz, ingredientBiz: ingredientBiz, Order: Order{MaxOrderCnt: conf.Rules.OrderMaxCnt}}
}

func (s *OrderService) OrderInfo(context.Context, *v1.GetOrderInfoReq) (*v1.GetOrderInfoReply, error) {
	var items []*v1.GetOrderInfoReply_Items
	for _, item := range s.items {
		items = append(items, &v1.GetOrderInfoReply_Items{
			Desc:  item.desc,
			Price: item.price,
		})
	}
	return &v1.GetOrderInfoReply{
		TotalAmount: s.GetPrice(),
		Items:       items,
	}, nil
}

func (s *OrderService) AddOrder(ctx context.Context, req *v1.AddOrderReq) (*v1.AddOrderReply, error) {
	if len(s.items) > int(s.MaxOrderCnt)-1 {
		return nil, ErrOrderOverflow
	}
	for _, order := range req.Items {
		ingredientCodes := order.GetIngredientCodes()
		item := NewOrderItem(s.productBiz, s.ingredientBiz)
		err := item.AddProduct(order.ProductCode, ingredientCodes...)
		if err != nil {
			return nil, err
		}
		s.items = append(s.items, *item)
	}

	return &v1.AddOrderReply{Result: true}, nil
}

type Order struct {
	items       []OrderItem
	MaxOrderCnt int32
}

func (o *Order) GetPrice() float32 {
	var totalAmount float32
	for _, item := range o.items {
		totalAmount += item.price
	}
	return totalAmount
}

func (o *Order) GetDesc() string {
	var desc string
	for i, item := range o.items {
		desc += fmt.Sprintf("%d. %s,价格：￥%0.2f\n", i+1, item.desc, item.price)
	}
	return desc
}
