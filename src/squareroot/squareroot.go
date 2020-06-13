package main

import (
	"math"
)

func squareroot() string  {
	x := 0.001
	for i := 0;  i<=1000000; i++ {
		x = math.Sqrt(x)
	}

	return "Code.education Rocks!"
}
