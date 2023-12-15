package service_test

import (
	"testing"

	"github.com/enrichman/coverage/service"
	"github.com/stretchr/testify/assert"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name string
		i, j int
		want int
	}{
		{name: "positive", i: 3, j: 5, want: 512},
		{name: "negative", i: -4, j: -9, want: -2197},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.Combine(tt.i, tt.j)
			assert.Equal(t, tt.want, result)
		})
	}
}
