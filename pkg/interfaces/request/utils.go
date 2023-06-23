package request

import (
	"math/rand"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/n-creativesystem/short-url/pkg/interfaces/response"
)

func validError(err error) *response.Errors {
	if err == nil {
		return nil
	}
	var errRes response.Errors
	if errs, ok := err.(validation.Errors); ok {
		for key, err := range errs {
			errs := response.ValidationErr2ErrorResponse(err, key)
			for _, e := range errs {
				errRes.Add(e)
			}
		}
	} else {
		errRes.AddError(err.Error(), "", "", "")
	}
	return &errRes
}

func randomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
