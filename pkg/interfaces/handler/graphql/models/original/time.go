package original

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/n-creativesystem/short-url/pkg/utils"
)

func MarshalTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		fmt.Fprintf(w, "%s", utils.TimeToString(t))
	})
}

func UnmarshalTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return utils.StringTimeParse(v)
	case []byte:
		return utils.UnmarshalBinaryTimeParse(v)
	default:
		return time.Time{}, fmt.Errorf("%T is not a time.Time", v)
	}
}
