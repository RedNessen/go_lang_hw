package main

import "fmt"

func main() {
	integer := 10
	floatPer := 9.9
	boolPer := true
	exception := "Hello"

	ptrInt := &integer
	ptrFloat := &floatPer
	ptrBool := &boolPer
	ptrString := &exception

	fmt.Println("Integer address - ", ptrInt)
	fmt.Println("Integer is - ", *ptrInt)
	fmt.Println("Float address - ", ptrFloat)
	fmt.Println("Float is - ", *ptrFloat)
	fmt.Println("Boolean address - ", ptrBool)
	fmt.Println("Boolean is - ", *ptrBool)
	fmt.Println("String address - ", ptrString)
	fmt.Println("String is - ", *ptrString)
}
