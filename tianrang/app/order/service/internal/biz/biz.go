package biz

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/pkg/errors"

	"log"
	"tianrang/app/order/service/internal/conf"
	"tianrang/app/order/service/internal/dao"
	V "tianrang/pkg/validate"
)

var ProviderSet = wire.NewSet(NewProductBiz, NewIngredientBiz)

type ProductRepo interface {
	LoadData(data []byte) []dao.Product
}

type IngredientRepo interface {
	LoadData(data []byte) []dao.Ingredient
}

type ProductBiz struct {
	repo ProductRepo

	products map[string]dao.Product
}

func NewProductBiz(repo ProductRepo) *ProductBiz {
	return &ProductBiz{repo: repo, products: make(map[string]dao.Product)}
}

func (biz *ProductBiz) LoadData(data []byte) {
	products := biz.repo.LoadData(data)
	for _, product := range products {
		biz.products[product.Code] = product
	}
}

func (biz *ProductBiz) GetProduct(code string) (dao.Product, bool) {
	v, ok := biz.products[code]
	return v, ok
}

type IngredientBiz struct {
	repo IngredientRepo

	ingredients map[string]dao.Ingredient
	rules       map[uint8]string
	validate    *validator.Validate
}

func NewIngredientBiz(repo IngredientRepo, rules *conf.Rules) *IngredientBiz {
	validate := validator.New()
	err := validate.RegisterValidation("cant_addition", V.CantAddition)
	if err != nil {
		log.Fatalf("add self validata：cant_addition err：%v\n", err)
	}
	err = validate.RegisterValidation("can_addition", V.CanAddition)
	if err != nil {
		log.Fatalf("add self validata：can_addition err：%v\n", err)
	}
	rulesMap := make(map[uint8]string)
	for _, rule := range rules.GetProductTypeRules() {
		rulesMap[uint8(rule.ProductType)] = rule.Rule
	}
	return &IngredientBiz{repo: repo, ingredients: make(map[string]dao.Ingredient), rules: rulesMap, validate: validate}
}

func (biz *IngredientBiz) LoadData(data []byte) {
	ingredients := biz.repo.LoadData(data)
	//i := make(map[string]Ingredient)
	for _, item := range ingredients {
		biz.ingredients[item.Code] = item
	}
}

func (biz *IngredientBiz) GetIngredients(codes ...string) []dao.Ingredient {
	var i []dao.Ingredient
	for _, code := range codes {
		if v, ok := biz.ingredients[code]; ok {
			i = append(i, v)
		}
	}

	return i
}

func (biz *IngredientBiz) CheckIngredients(productType uint8, codes ...string) error {

	rules := map[string]interface{}{
		"ingredient_codes": biz.rules[productType],
	}
	data := map[string]interface{}{
		"ingredient_codes": codes,
	}
	errs := biz.validate.ValidateMap(data, rules)
	if len(errs) > 0 {
		// 记录错误日志，或者返回err
		for _, err := range errs {
			errs := err.(validator.ValidationErrors)
			for _, e := range errs {
				// can translate each error one at a time.
				//fmt.Println(e.Error())
				return errors.New(e.Error())
			}
		}
	}
	return nil
}
