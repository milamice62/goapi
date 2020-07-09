package models

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func validationStruct(s interface{}) []string {
	validate = validator.New()
	err := validate.Struct(s)
	if err != nil {
		fieldErrors := []string{}
		for _, fe := range err.(validator.ValidationErrors) {
			fieldErrors = append(fieldErrors, fmt.Sprintf("error: %s should be %s", fe.Field(), fe.Tag()))
		}
		return fieldErrors
	}
	return nil
}
