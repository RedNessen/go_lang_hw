package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("App shut down")
	}()

	fmt.Println("App started")
	open()
	fmt.Println("Working...")
	fmt.Println("Tasks completed")
	close()

}

func close() {
	defer func() {
		fmt.Println("Connection is closed")
	}()
	fmt.Println("Function close is called")
}

func open() {
	defer func() {
		fmt.Println("Connection is opened")
	}()
	fmt.Println("App is trying to open connect")
}
