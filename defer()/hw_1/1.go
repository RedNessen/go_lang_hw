package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Function is complete")
	}()

	defer func() {
		fmt.Println("DB connection is closed")
	}()

	fmt.Println("Program is working")
	fmt.Println("Program is opened DB connection")
	fmt.Println("Task is complete")
}
