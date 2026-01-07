package main

import "fmt"

type Car struct {
	Type string

	Gasoline bool

	Doors int
}

func (car *Car) changeDoorsNumber(doorsNumber int) {
	if doorsNumber > 7 {
		fmt.Println("Validation on doors is failed, result is greater than 7")
	}
}

func (car Car) upgradeTypePreview() {
	if car.Gasoline {
		car.Gasoline = false
		fmt.Println("This is how your vechicle could look after upgrade - ", car)
	} else {
		fmt.Println("Your car already electric)")
	}
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
	toyotaPrius.changeDoorsNumber(3)
	toyotaPrius.upgradeTypePreview()
	fmt.Println(toyotaPrius)
	fmt.Println()

	fmt.Println(tesla)
	tesla.changeDoorsNumber(8)
	tesla.upgradeTypePreview()
	fmt.Println(tesla)

}
