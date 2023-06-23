package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/n-creativesystem/short-url/pkg/interfaces/response"
)

type RegisterApplication struct {
	// UserId  string `json:"user_id"`
	AppName string `json:"application_name"`
}

func (r *RegisterApplication) Valid() *response.Errors {
	err := validation.ValidateStruct(r,
		validation.Field(
			&r.AppName,
			validation.Required.Error("application_name is a required field."),
			validation.When(r.AppName != "", validation.Length(1, 255).Error("application_name is a {{.min}} ~ {{.max}} characters.")),
		),
	)
	return validError(err)
}
