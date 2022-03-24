package dao

import "encoding/json"

// Product 产品
type Product struct {
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Price float32 `json:"price"`
	Type  uint8   `json:"type"`
}


type productRepo struct {
}

func NewProductRepo() *productRepo {
	return &productRepo{}
}

func (repo *productRepo) LoadData(data []byte) []Product {
	var p []Product
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil
	}
	return p
}