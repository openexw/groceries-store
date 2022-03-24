package validate

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

// data := make(map[string]string)
// rules "= map[string]string{"ingredients"}

func CantAddition(fl validator.FieldLevel) bool {
	val := fl.Field().Interface()
	params := parseCantAdditionParams(fl.Param())
	for _, param := range params {
		item := strings.Split(param, " ")
		if reflect.DeepEqual(item, val) {
			return false
		}
	}
	return true
}

// parseCantAdditionParams s="a b;c d"
func parseCantAdditionParams(s string) []string {
	return strings.Split(s, ";")
}
