package main

import "fmt"

func main() {
	breastSize := 2
	sizePtr := &breastSize

	sizeChange(sizePtr, 3)

}

func sizeChange(sizePtr *int, desiredSize int) {
	if sizePtr != nil {
		fmt.Println("Operation has began, current size = ", *sizePtr)
		*sizePtr = desiredSize
		fmt.Println("Operation successful, new size = ", *sizePtr)
	} else {
		fmt.Println("Error. Pointer is nil")
	}
}
