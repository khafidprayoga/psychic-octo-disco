package utils

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

func ValidateStruct[T any](input T) error {
	enNew := en.New()
	translator := ut.New(enNew)
	validate := validator.New()
	trans, _ := translator.GetTranslator("en_US")
	enTranslations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Translate(trans))
		}
	}
	return nil
}
