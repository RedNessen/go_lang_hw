package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Cube of digit 3.14 = ", floatInCube(3.14))
}

func floatInCube(x float64) float64 {
	return math.Pow(x, 3)
}
