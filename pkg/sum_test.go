package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	test := []struct {
		name   string
		values []int
		want   int
	}{
		{
			name:   "simple test",
			values: []int{1, 2},
			want:   3,
		},
		{
			name:   "one",
			values: []int{1},
			want:   1,
		},
		{
			name:   "with negative values",
			values: []int{-1, -2, 3},
			want:   0,
		},
	}

	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, Sum(test.values...), test.want)
		})
	}
}
