package original

import (
	"fmt"
	"io"
	"net/url"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalURL(u url.URL) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		fmt.Fprintf(w, "%s", u.String())
	})
}

func UnmarshalURL(v interface{}) (url.URL, error) {
	switch v := v.(type) {
	case string:
		result, err := url.Parse(v)
		if err != nil {
			return url.URL{}, err
		}
		return *result, nil
	case []byte:
		result := &url.URL{}
		if err := result.UnmarshalBinary(v); err != nil {
			return *result, err
		}
		return url.URL{}, nil
	default:
		return url.URL{}, fmt.Errorf("%T is not a url.URL", v)
	}
}
