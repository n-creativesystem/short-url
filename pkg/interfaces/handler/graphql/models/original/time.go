package original

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/n-creativesystem/short-url/pkg/utils"
)

func MarshalTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		v := utils.TimeToString(t)
		_ = json.NewEncoder(w).Encode(v)
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
