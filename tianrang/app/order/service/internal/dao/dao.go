package dao

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewIngredientRepo, NewProductRepo)
