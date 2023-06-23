package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	result := SplitArray(arr, 7)
	assert.Equal(t, 3, len(result))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, result[0])
	assert.Equal(t, []int{8, 9, 10, 11, 12, 13, 14}, result[1])
	assert.Equal(t, []int{15, 16, 17, 18, 19, 20}, result[2])

	result = SplitArray(arr, 0)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, result[0])
	assert.Equal(t, []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, result[1])
}
