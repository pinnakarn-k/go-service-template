package validator

import (
	"errors"
	"fmt"

	playgroundvalidator "github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *playgroundvalidator.Validate
}

func New() *Validator {
	return &Validator{
		validate: playgroundvalidator.New(),
	}
}

func (v *Validator) Validate(input any) error {
	err := v.validate.Struct(input)
	if err == nil {
		return nil
	}

	var validationErrors playgroundvalidator.ValidationErrors
	if errors.As(err, &validationErrors) {
		details := mapValidationErrors(validationErrors)

		return newValidationError(details...)
	}

	return fmt.Errorf("unexpected validator error: %w", err)
}
