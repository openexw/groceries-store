package validate

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func CanAddition(fl validator.FieldLevel) bool {
	vals := fl.Field().Interface()
	params := strings.Split(fl.Param(), " ")
	paramsMap := make(map[string]struct{})
	for _, param := range params {
		paramsMap[param] = struct{}{}
	}

	for _, v := range vals.([]string) {
		if _, ok := paramsMap[v]; !ok {
			return false
		}
	}

	return true
}

func xx(params, vals []string) {
	paramsMap := make(map[string]struct{})
	for _, param := range params {
		paramsMap[param] = struct{}{}
	}

	isE := true
	for _, v := range vals {
		if _, ok := paramsMap[v]; !ok {
			isE = false
			break
		}
	}
	println(isE)
}
