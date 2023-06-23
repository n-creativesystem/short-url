package short

import validation "github.com/go-ozzo/ozzo-validation/v4"

type ValidationError validation.Error

type ValidationErrors struct {
	validation.Errors
}
