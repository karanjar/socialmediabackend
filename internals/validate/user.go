package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Users(s interface{}) error {
	err := _validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); !ok {
			panic(err)

		}

		for _, errs := range err.(validator.ValidationErrors) {
			fmt.Println("Validation Error:", errs.Field())
			fmt.Println(errs.Type())
		}

		return err
	}
	return nil

}
