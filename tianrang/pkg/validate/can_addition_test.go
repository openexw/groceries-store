package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"testing"
)

func TestCanAddition(t *testing.T) {
	type ingredient struct {
		IngredientCodes []string `validate:"max=1,can_addition=a b c d"`
	}

	validate := validator.New()
	err := validate.RegisterValidation("can_addition", CanAddition)
	if err != nil {
		log.Fatalln(err)
	}

	i := ingredient{IngredientCodes: []string{"e"}}
	if err := validate.Struct(i); err != nil {

		for _, e := range err.(validator.ValidationErrors) {
			fmt.Printf("%v", e.Error())
		}
	}
}

func Test_xx(t *testing.T) {
	xx([]string{"a", "b", "c", "d"}, []string{"a"})
	xx([]string{"a", "b", "c", "d"}, []string{"a", "b"})
	xx([]string{"a", "b", "c", "d"}, []string{"a", "h"})
	xx([]string{"a", "b", "c", "d"}, []string{"a", "d"})
	xx([]string{"a", "b", "c", "d"}, []string{"c", "d"})
	xx([]string{"a", "b", "c", "d"}, []string{""})
}
