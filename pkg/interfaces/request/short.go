package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/n-creativesystem/short-url/pkg/interfaces/response"
)

type GenerateShortURL struct {
	URL string `json:"url"`
	Key string `json:"key,omitempty"` //nullable
}

func (req *GenerateShortURL) Valid() *response.Errors {
	err := validation.ValidateStruct(req,
		validation.Field(
			&req.URL,
			validation.Required.Error("URL is a required field."),
			is.URL.Error("Please enter in URL format."),
		),
		validation.Field(
			&req.Key,
			validation.When(req.Key != "",
				validation.RuneLength(1, 255).Error("The key range is {{.min}} to {{.max}}."),
			),
		),
	)
	return validError(err)
}

type RequestPathForGenerateQRCode struct {
	Key string
}

func (req *RequestPathForGenerateQRCode) Valid() *response.Errors {
	err := validation.ValidateStruct(req,
		validation.Field(
			&req.Key,
			validation.RuneLength(1, 255).Error("The key range is {{.min}} to {{.max}}."),
		),
	)
	return validError(err)
}
