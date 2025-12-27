package main

import "fmt"

func main() {
	fmt.Println("You`re met a friend and tell him:")
	privet()
	fmt.Println("You`re leaving and saying:")
	poka()
	fmt.Println("You`re home and saying to your brother:")
	privet()
}

func privet() {
	fmt.Println("Privet!)")
}

func poka() {
	fmt.Println("Poka!(")
}
