package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

// testStruct 匿名構造体でのジェネリクスは許可されていないので一旦構造体を定義する
type testStruct[T constraints.Ordered] struct {
	name    string
	value   Required[T]
	want    T
	wantErr bool
}

func TestRequiredByString(t *testing.T) {
	tests := []testStruct[string]{
		{
			name:    "valid",
			value:   String("aaa"),
			want:    "aaa",
			wantErr: false,
		},
		{
			name:    "invalid",
			value:   String(""),
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := tt.value.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Required() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				assert.Equal(t, ErrInvalid, err)
			}
			assert.Equal(t, tt.want, v)
		})
	}
}

func TestRequiredByNumeric(t *testing.T) {
	tests := []testStruct[float64]{
		{
			name:    "valid",
			value:   Numeric(1.0),
			want:    1.0,
			wantErr: false,
		},
		{
			name:    "invalid",
			value:   Numeric(0.0),
			want:    0.0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := tt.value.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Required() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				assert.Equal(t, ErrInvalid, err)
			}
			assert.Equal(t, tt.want, v)
		})
	}
}
