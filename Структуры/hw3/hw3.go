package main

import "fmt"

type Appartment struct {
	Address     string
	SquareSize  float64
	FloorNumber int
	Price       int
}

func NewAppartment(
	address string,
	squareSize float64,
	floorNumber int,
	price int,
) Appartment {
	if address == "" {
		return Appartment{}
	}

	if squareSize < 10.0 {
		return Appartment{}
	}

	if floorNumber < 1 || floorNumber > 100 {
		return Appartment{}
	}

	if price < 1 {
		return Appartment{}
	}

	return Appartment{
		Address:     address,
		SquareSize:  squareSize,
		FloorNumber: floorNumber,
		Price:       price,
	}
}

func (appartment *Appartment) changePrice(newPrice int) {
	if newPrice > 0 {
		appartment.Price = newPrice
	} else {
		fmt.Println("Price is less then a zero")
	}
}

func main() {

	appartment1 := NewAppartment("Lenina 39", 44.3, 4, 4883729)
	appartment1ptr := &appartment1

	fmt.Println(*appartment1ptr)

	appartment2 := NewAppartment("MKAD 32", 10.0, 2, 312342123)
	appartment2ptr := &appartment2

	fmt.Println(*appartment2ptr)

	appartment2.changePrice(0)
	fmt.Println(*appartment2ptr)

}
