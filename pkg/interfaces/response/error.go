package response

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Error struct {
	Message     string `json:"message"`
	Fields      string `json:"field"`
	Help        string `json:"help"`
	Description string `json:"description,omitempty"`
}

func (e *Error) SetFieldWhenNil(field string) {
	if e.Fields == "" {
		e.Fields = field
	}
}

func (e Error) Error() string {
	return e.Message
}

func NewError(message, field, help, description string) Error {
	return Error{
		message,
		field,
		help,
		description,
	}
}

type Errors struct {
	Errors []Error `json:"errors"`
}

func (e *Errors) AddError(message, field, help, description string) {
	e.Add(NewError(message, field, help, description))
}

func (e *Errors) Add(err Error) {
	e.Errors = append(e.Errors, err)
}

func (e *Errors) Error() string {
	var err error
	for _, er := range e.Errors {
		err = errors.Join(err, er)
	}
	return err.Error()
}

func (e *Errors) GraphQLError(ctx context.Context) error {
	errList := gqlerror.List{}
	for _, er := range e.Errors {
		extensions := Extensions{}.SetCode(http.StatusBadRequest)
		errList = append(errList, &gqlerror.Error{
			Path:       graphql.GetPath(ctx),
			Message:    er.Message,
			Extensions: extensions,
		})
	}
	return errList
}

func NewErrors(err error) *Errors {
	return NewErrorsWithMessage(err.Error())
}

func NewErrorsWithMessage(message string) *Errors {
	return &Errors{
		Errors: []Error{
			NewError(message, "", "", ""),
		},
	}
}

func ValidationErr2ErrorResponse(err error, key string) []Error {
	results := make([]Error, 0, 10)

	switch er := err.(type) {
	case validation.Errors:
		for k, e := range er {
			results = append(results, ValidationErr2ErrorResponse(e, fmt.Sprintf("%s.%s", key, k))...)
		}
	case validation.Error:
		params := er.Params()
		results = append(results, Error{
			Message: htmlTemplate(er.Message(), params),
			Fields:  key,
			Help:    "",
		})
	default:
		results = append(results, Error{
			Message: err.Error(),
			Fields:  key,
			Help:    "",
		})
	}

	return results
}

func htmlTemplate(msg string, params map[string]interface{}) string {
	if params == nil {
		params = make(map[string]interface{})
	}
	buf := new(bytes.Buffer)
	tpl, err := template.New("msg").Parse(msg)
	if err != nil {
		return msg
	}
	err = tpl.Execute(buf, params)
	if err != nil {
		return msg
	}
	return buf.String()
}
