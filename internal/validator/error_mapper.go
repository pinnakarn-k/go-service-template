package validator

import (
	"fmt"
	"go-service-template/internal/apperror"
	"strings"

	playgroundvalidator "github.com/go-playground/validator/v10"
)

func mapValidationErrors(
	errs playgroundvalidator.ValidationErrors,
) []apperror.Detail {
	details := make([]apperror.Detail, 0, len(errs))

	for _, fieldErr := range errs {
		details = append(details, apperror.Detail{
			Field:   formatFieldName(fieldErr.Field()),
			Message: validationMessage(fieldErr),
		})
	}

	return details
}

func validationMessage(
	fieldErr playgroundvalidator.FieldError,
) string {
	switch fieldErr.Tag() {
	case "required":
		return "is required"

	case "min":
		return fmt.Sprintf(
			"must be greater than or equal to %s",
			fieldErr.Param(),
		)

	case "max":
		return fmt.Sprintf(
			"must be less than or equal to %s",
			fieldErr.Param(),
		)

	case "gt":
		return fmt.Sprintf(
			"must be greater than %s",
			fieldErr.Param(),
		)

	case "gte":
		return fmt.Sprintf(
			"must be greater than or equal to %s",
			fieldErr.Param(),
		)

	case "lt":
		return fmt.Sprintf(
			"must be less than %s",
			fieldErr.Param(),
		)

	case "lte":
		return fmt.Sprintf(
			"must be less than or equal to %s",
			fieldErr.Param(),
		)

	case "oneof":
		return fmt.Sprintf(
			"must be one of: %s",
			fieldErr.Param(),
		)

	case "email":
		return "must be a valid email address"

	case "uuid":
		return "must be a valid UUID"

	case "datetime":
		return "must be a valid date in YYYY-MM-DD format"

	default:
		return fmt.Sprintf(
			"failed validation rule %s",
			fieldErr.Tag(),
		)
	}
}

func formatFieldName(field string) string {
	if field == "" {
		return ""
	}

	return strings.ToLower(field[:1]) + field[1:]
}
