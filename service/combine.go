package service

import (
	"github.com/enrichman/coverage/math_01"
	"github.com/enrichman/coverage/math_02"
	"github.com/enrichman/coverage/math_03"
)

func Combine(i, j int) int {
	sum1 := math_01.Sum(i, j)
	sum2 := math_02.Sum(i, j)
	sum3 := math_03.Sum(i, j)

	mult2 := math_02.Multiply(sum1, sum2)
	mult3 := math_03.Multiply(sum3, mult2)

	return mult3
}
