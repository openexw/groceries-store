package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"testing"
)

func TestCantAddition(t *testing.T) {
	type ingredient struct {
		IngredientCodes []string `validate:"max=1,oneof=oreo_milk_cover cheese_milk_cover"`
	}

	validate := validator.New()
	err := validate.RegisterValidation("cant_addition", CanAddition)
	if err != nil {
		log.Fatalln(err)
	}

	i := ingredient{IngredientCodes: []string{"a", "b"}}
	if err := validate.Struct(i); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			println(e.Value())
		}
	}
}

func TestCantAdditionMap(t *testing.T) {
	rules := map[string]interface{}{
		//"IngredientCodes": "cant_addition=a c;d c",
		//"ingredient_codes": "max=1,oneof=oreo_milk_cover cheese_milk_cover,cant_addition=oreo_milk_cover cheese_milk_cover",
		"ingredient_codes": "max=1,oneof=oreo_milk_cover cheese_milk_cover",
		//"ingredient_codes": "max=1,cant_addition=oreo_milk_cover cheese_milk_cover",
	}
	data := map[string]interface{}{
		"ingredient_codes": []string{"cheese_milk_cover", "a"},
	}
	validate := validator.New()
	err := validate.RegisterValidation("cant_addition", CantAddition)
	if err != nil {
		log.Fatalln(err)
	}
	errs := validate.ValidateMap(data, rules)
	if len(errs) > 0 {
		fmt.Printf("%v", errs)
	}
}
