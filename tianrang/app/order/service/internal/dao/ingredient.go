package dao

import (
	"encoding/json"
)


// Ingredient 配料
type Ingredient struct {
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Price float32 `json:"price"`
}

type ingredientRepo struct {
}

func NewIngredientRepo() *ingredientRepo {
	return &ingredientRepo{}
}

func (repo ingredientRepo) LoadData(data []byte) []Ingredient {
	var i []Ingredient
	err := json.Unmarshal(data, &i)
	if err != nil {
		return nil
	}
	return i
}
