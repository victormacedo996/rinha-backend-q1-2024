package validator

import (
	"sync"

	playValidator "github.com/go-playground/validator/v10"
)

var once sync.Once

var Validator *playValidator.Validate

func validateCreditOrDebit(fl playValidator.FieldLevel) bool {
	r, ok := fl.Field().Interface().(rune)
	if !ok {
		return false
	}

	return r == 'c' || r == 'd'
}

func GetInstance() *playValidator.Validate {
	if Validator == nil {
		once.Do(
			func() {
				Validator = playValidator.New()
				Validator.RegisterValidation("validateCreditOrDebit", validateCreditOrDebit)

			},
		)
	}

	return Validator

}
