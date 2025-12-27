package main

import "fmt"

var globalTemp int = 15

func main() {
	fmt.Println("Temperature = ", globalTemp)
	tempGoColder()
	fmt.Println("Temperature = ", globalTemp)
	tempGoColder()
	fmt.Println("Temperature = ", globalTemp)
	tempGoHiger()
	fmt.Println("Temperature = ", globalTemp)
	tempGoColder()
	tempGoColder()
	tempGoColder()
	fmt.Println("Temperature = ", globalTemp)
}

func tempGoColder() {
	globalTemp -= 5
}

func tempGoHiger() {
	globalTemp += 5
}
