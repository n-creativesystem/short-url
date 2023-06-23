package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCSRFConfig(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Csrf
		want    string
		wantErr bool
	}{
		{
			name:    "Required.",
			cfg:     Csrf{},
			want:    ErrCSRFSetting.Error(),
			wantErr: true,
		},
		{
			name: "Missing Header.",
			cfg: Csrf{
				TokenBase:             false,
				CorsAndOriginalHeader: true,
			},
			want:    "HeaderName: cannot be blank.",
			wantErr: true,
		},
		{
			name: "Missing Header name.",
			cfg: Csrf{
				TokenBase:             false,
				CorsAndOriginalHeader: true,
				HeaderName:            "HeaderName",
			},
			want:    "HeaderName: Start with `X-` or `x-`.",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			err := tt.cfg.IsValid()
			if tt.wantErr {
				require.Error(err)
				require.EqualError(err, tt.want)
			}
		})
	}
}
