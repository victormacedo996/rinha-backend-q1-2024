package validator

import (
	"sync"

	playValidator "github.com/go-playground/validator/v10"
)

var once sync.Once

var validator *playValidator.Validate

func validateCreditOrDebit(fl playValidator.FieldLevel) bool {
	r, ok := fl.Field().Interface().(rune)
	if !ok {
		return false
	}

	return r == 'c' || r == 'd'
}

func GetInstance() *playValidator.Validate {
	if validator == nil {
		once.Do(
			func() {
				validator = playValidator.New()
				validator.RegisterValidation("validateCreditOrDebit", validateCreditOrDebit)

			},
		)
	}

	return validator

}
