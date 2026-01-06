package main

import "fmt"

func main() {
	xInt := 33
	xFloat := 44.4
	xBool := true
	xString := "Test string"

	ptrInt := &xInt
	ptrFloat := &xFloat
	ptrBool := &xBool
	ptrString := &xString

	fmt.Println("Integer")
	fmt.Println(*ptrInt)
	fmt.Println("Func has been called")
	intChange(ptrInt)
	fmt.Println(*ptrInt)
	fmt.Println("")

	fmt.Println("Float")
	fmt.Println(*ptrFloat)
	fmt.Println("Func has been called")
	floatChange(ptrFloat)
	fmt.Println(*ptrFloat)
	fmt.Println("")

	fmt.Println("Bool")
	fmt.Println(*ptrBool)
	fmt.Println("Func has been called")
	boolChange(ptrBool)
	fmt.Println(*ptrBool)
	fmt.Println("")

	fmt.Println("String")
	fmt.Println(*ptrString)
	fmt.Println("Func has been called")
	stringChange(ptrString)
	fmt.Println(*ptrString)

}

func intChange(x *int) {
	*x += 11
}

func floatChange(x *float64) {
	*x += 11.1
}

func boolChange(x *bool) {
	*x = !*x
}

func stringChange(x *string) {
	*x = "Func has been completed"
}
