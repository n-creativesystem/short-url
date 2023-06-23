package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	type testTable struct {
		name    string
		baseURL string
		paths   []string
		want    string
		wantErr error
	}
	tests := []testTable{
		{
			name:    "success",
			baseURL: "http://localhost",
			paths:   []string{"test", "about"},
			want:    "http://localhost/test/about",
			wantErr: nil,
		},
		{
			name:    "failed parse",
			baseURL: "http://192.168.0.%31/",
			paths:   nil,
			want:    "",
			wantErr: errors.New(`parse "http://192.168.0.%31/": invalid URL escape "%31"`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := URL(tt.baseURL, tt.paths...)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("URL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, u)
		})
	}
}

func TestMustURL(t *testing.T) {
	u := MustURL("http://localhost", "test", "about")
	assert.Equal(t, "http://localhost/test/about", u)
	defer func() {
		err := recover().(error)
		assert.Equal(t, `parse "http://192.168.0.%31/": invalid URL escape "%31"`, err.Error())
	}()
	_ = MustURL("http://192.168.0.%31/", "test")
}
