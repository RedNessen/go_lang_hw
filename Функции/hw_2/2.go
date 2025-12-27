package main

import (
	"fmt"
	"time"
)

func main() {
	greetGuest("Alexander", 401)
	greetGuest("Anastasiy", 321)
	greetGuest("Ilya", 112)
}

func greetGuest(name string, roomNumber int) {
	fmt.Println("Greetings, Mr", name, "we`re welcome you in our hotel!")
	fmt.Println("Please, i`m going to escort you in room number", roomNumber)
	fmt.Println()
	time.Sleep(500 * time.Millisecond)
}
