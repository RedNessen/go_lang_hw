package main

import "fmt"

func main() {
	var liters float64 = 3.0

	if liters > 3.0 || liters < 3.0 {
		fmt.Println("Ты какой-то странный...")
	} else {
		fmt.Println("А ты знаешь золотую середину!")
	}
}
