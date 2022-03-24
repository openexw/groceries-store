package service

import (
	"github.com/pkg/errors"
	"tianrang/app/order/service/internal/biz"
)

type OrderItem struct {
	productBiz    *biz.ProductBiz
	ingredientBiz *biz.IngredientBiz

	MaxIngredientCnt uint8
	ingredients      []string
	productCode      string
	price            float32
	desc             string
}

func NewOrderItem(productBiz *biz.ProductBiz, ingredientBiz *biz.IngredientBiz) *OrderItem {
	return &OrderItem{productBiz: productBiz, ingredientBiz: ingredientBiz}
}

func (order *OrderItem) AddProduct(productCode string, ingredients ...string) error {
	info, ok := order.productBiz.GetProduct(productCode)
	if !ok {
		return ErrProductNotFound
	}

	order.price = info.Price
	order.productCode = productCode
	order.desc = info.Name

	// 没有添加配料无需下面的逻辑
	if len(ingredients) <= 0 {
		return nil
	}
	ingredientList := order.ingredientBiz.GetIngredients(ingredients...)
	if len(ingredientList) <= 0 {
		return ErrIngredientsNotFound
	}
	// check the rules
	if err := order.ingredientBiz.CheckIngredients(info.Type, ingredients...); err != nil {
		return errors.Wrap(err, "添加配料错误")
	}

	var ingredientStr string
	for _, ingredient := range ingredientList {
		order.ingredients = append(order.ingredients, ingredient.Code)
		order.price += ingredient.Price
		ingredientStr += ingredient.Name + "、"
	}
	order.desc += "\n配料：" + ingredientStr[:len(ingredientStr)-3]
	return nil
}
