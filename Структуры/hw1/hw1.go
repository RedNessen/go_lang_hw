package main

import "fmt"

type Car struct {
	Type string

	Gasoline bool

	Doors int
}

func main() {

	toyotaPrius := Car{}
	toyotaPrius.Type = "Sedan"
	toyotaPrius.Gasoline = true
	toyotaPrius.Doors = 5

	tesla := Car{}
	tesla.Type = "Crossover"
	tesla.Gasoline = false
	tesla.Doors = 5

	fmt.Println(toyotaPrius)
	toyotaPrius.Gasoline = false
	fmt.Println(toyotaPrius)
	fmt.Println()

	fmt.Println(tesla)
	tesla.Type = "Sedan"
	fmt.Println(tesla)

}
