package math_03_test

import (
	"testing"

	math "github.com/enrichman/coverage/math_03"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		i, j int
		want int
	}{
		{name: "positive", i: 3, j: 5, want: 8},
		{name: "positive 2", i: 1, j: 15, want: 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := math.Sum(tt.i, tt.j)
			assert.Equal(t, tt.want, sum)
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		i, j int
		want int
	}{
		{name: "simple mult", i: 3, j: 5, want: 15},
		{name: "one negative", i: -2, j: 4, want: -8},
		{name: "both negative", i: -5, j: -9, want: 45},
		{name: "zero", i: 5, j: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := math.Multiply(tt.i, tt.j)
			assert.Equal(t, tt.want, sum)
		})
	}
}
