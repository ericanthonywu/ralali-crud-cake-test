package Utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"ralali-crud-cake-test/Model"
)

func Validate(request interface{}) []Model.ValidationError {
	var v = validator.New()
	if errs := v.Struct(request); errs != nil {
		var validationErrors []Model.ValidationError
		for _, validationError := range errs.(validator.ValidationErrors) {
			var elem Model.ValidationError

			elem.Message = fmt.Sprintf("%s is %s", validationError.Field(), validationError.Tag())
			elem.Value = validationError.Value()

			validationErrors = append(validationErrors, elem)
		}
		return validationErrors
	}
	return nil
}
