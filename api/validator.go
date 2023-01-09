package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/rajatk-31/simplebank/util"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		//check if currency is valid or not
		return util.IsSupportedCurrency(currency)
	}
	return false
}
