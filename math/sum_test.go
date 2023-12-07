package math_test

import (
	"testing"

	"github.com/enrichman/coverage/math"
)

func TestSum(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive",
			args: args{i: 3, j: 5},
			want: 8,
		},

		{
			name: "negative",
			args: args{i: 3, j: -5},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := math.Sum(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
