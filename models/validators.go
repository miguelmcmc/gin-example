package models

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

var (
	errInvalidType      = errors.New("Error invalid struct type").Error()
	errInvalidStructure = errors.New("Error invalid struct format").Error()
)

//ValidateStruct s with validate fields
func ValidateStruct(s interface{}) error {
	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(s)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return errors.Wrap(err, errInvalidType)
		}

		// for _, err := range err.(validator.ValidationErrors) {

		// fmt.Println(err.Namespace())
		// fmt.Println(err.Field())
		// fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
		// fmt.Println(err.StructField())     // by passing alt name to ReportError like below
		// fmt.Println(err.Tag())
		// fmt.Println(err.ActualTag())
		// fmt.Println(err.Kind())
		// fmt.Println(err.Type())
		// fmt.Println(err.Value())
		// fmt.Println(err.Param())

		// return errors.New("Error in: " + err.StructField())
		// }

		// from here you can create your own error messages in whatever language you wish
		return errors.Wrap(err, errInvalidStructure)
	}

	return nil
}
