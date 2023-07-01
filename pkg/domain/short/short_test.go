package short

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShortsWithTimestamps(t *testing.T) {
	values := []ShortWithTimeStamp{
		{
			Short: &Short{
				operation: none,
			},
		},
		{
			Short: &Short{
				operation: none,
			},
		},
	}
	v := ShortWithTimeStamps(values)
	v.New()
	for _, value := range v {
		require.Equal(t, value.operation, new)
	}
	v.Update()
	for _, value := range v {
		require.Equal(t, value.operation, update)
	}
}
