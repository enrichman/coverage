package main

import (
	"os"
	"strconv"

	"github.com/enrichman/coverage/math"
)

func main() {
	if len(os.Args) > 2 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			os.Exit(1)
		}

		j, err := strconv.Atoi(os.Args[2])
		if err != nil {
			os.Exit(1)
		}

		math.Sum(i, j)
	}
}
