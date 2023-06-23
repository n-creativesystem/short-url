package request

import (
	"testing"

	"github.com/n-creativesystem/short-url/pkg/interfaces/response"
	"github.com/stretchr/testify/require"
)

func TestOAuth2ClientValid(t *testing.T) {
	var err *response.Errors
	r := require.New(t)
	client := RegisterApplication{}
	client.AppName = randomString(1)
	err = client.Valid()
	r.Nil(err)

	client.AppName = randomString(255)
	err = client.Valid()
	r.Nil(err)

	client.AppName = "app_name"
	err = client.Valid()
	r.Nil(err)

	client.AppName = ""
	err = client.Valid()
	r.NotNil(err)
	r.Len(err.Errors, 1)
	r.Equal(err.Errors[0].Fields, "application_name")
	r.Equal(err.Errors[0].Message, "application_name is a required field.")

	client.AppName = randomString(256)
	err = client.Valid()
	r.NotNil(err)
	r.Len(err.Errors, 1)
	r.Equal(err.Errors[0].Fields, "application_name")
	r.Equal(err.Errors[0].Message, "application_name is a 1 ~ 255 characters.")
}
