package main

import "fmt"

func main() {
	var beerDelicious = false
	var croutonsDelicious = false

	if beerDelicious && croutonsDelicious {
		fmt.Println("Мы гулять идем!")
	} else {
		fmt.Println("Мы гулять не идем!")
	}
}
