package validation

import (
	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func init() {
	validation.ErrRequired = validation.ErrRequired.SetMessage("the value is required")
	validation.ErrLengthOutOfRange = validation.ErrLengthOutOfRange.SetMessage("the value is {{.min}} ~ {{.max}} characters")
	var ErrURL = is.ErrURL.SetMessage("the value is valid URL")
	is.URL = validation.NewStringRuleWithError(govalidator.IsURL, ErrURL)
}
