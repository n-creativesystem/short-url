package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertContentTypeToEnum(t *testing.T) {
	type testTable struct {
		name        string
		contentType string
		want        contentType
	}

	tests := []testTable{
		{
			name:        "application/json",
			contentType: "application/json",
			want:        jsonType,
		},
		{
			name:        "application/json and charset",
			contentType: "application/json;charset=UTF-8",
			want:        jsonType,
		},
		{
			name:        "application/x-www-form-urlencoded",
			contentType: "application/x-www-form-urlencoded",
			want:        formType,
		},
		{
			name:        "Other",
			contentType: "application/xml",
			want:        otherType,
		},
	}
	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			typ := convertContentTypeToEnum(tt.contentType)
			assert.Equal(t, tt.want, typ)

			r, err := http.NewRequest(http.MethodPost, "http://localhost", nil)
			require.NoError(t, err)
			r.Header.Set("content-type", tt.contentType)
			typ = convertContentTypeToEnumWithRequest(r)
			assert.Equal(t, tt.want, typ)
		})
	}
}
