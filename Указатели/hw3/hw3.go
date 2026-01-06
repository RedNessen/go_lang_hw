package main

import "fmt"

func main() {
	xInt := 33
	// xFloat := 44.4
	// xBool := true
	xString := "Test string"

	ptrInt := &xInt
	// ptrFloat := &xFloat
	// ptrBool := &xBool
	ptrString := &xString

	// var ptrIntNil *int
	var ptrFloatNil *float64
	var ptrBoolNil *bool
	// var ptrStringNil *string

	printInt(ptrInt)
	printFloat(ptrFloatNil)
	printBool(ptrBoolNil)
	printString(ptrString)
}

func printInt(intPtr *int) {
	if intPtr != nil {
		fmt.Println(*intPtr)
	} else {
		fmt.Println("Pointer is nil")
	}
}

func printFloat(floatPtr *float64) {
	if floatPtr != nil {
		fmt.Println(*floatPtr)
	} else {
		fmt.Println("Pointer is nil")
	}
}

func printBool(boolPtr *bool) {
	if boolPtr != nil {
		fmt.Println(*boolPtr)
	} else {
		fmt.Println("Pointer is nil")
	}
}

func printString(stringPtr *string) {
	if stringPtr != nil {
		fmt.Println(*stringPtr)
	} else {
		fmt.Println("Pointer is nil")
	}
}
