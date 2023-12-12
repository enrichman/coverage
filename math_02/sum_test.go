package math_02_test

import (
	"testing"

	math "github.com/enrichman/coverage/math_02"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		i, j int
		want int
	}{
		{
			name: "positive",
			i:    3,
			j:    5,
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := math.Sum(tt.i, tt.j)
			assert.Equal(t, tt.want, sum)
		})
	}
}
