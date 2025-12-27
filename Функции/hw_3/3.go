package main

import (
	"fmt"
	"math"
)

func main() {
	QuadraticRootsInt(1, -3, 2)
}

func QuadraticRootsInt(a, b, c int) {
	if a == 0 {
		fmt.Println("No result")
		return
	}

	d := b*b - 4*a*c
	if d < 0 {
		fmt.Println("No result")
		return
	}

	sqrtD := int(math.Sqrt(float64(d)))
	if sqrtD*sqrtD != d {
		fmt.Println("No result")
		return
	}

	den := 2 * a
	if den == 0 {
		fmt.Println("No result")
		return
	}

	if (-b+sqrtD)%den != 0 || (-b-sqrtD)%den != 0 {
		fmt.Println("No result")
		return
	}

	fmt.Println((-b + sqrtD) / den)
	fmt.Println((-b - sqrtD) / den)
}
