package main

import "github.com/k0kubun/pp"

type car struct {
	Type  string
	Doors int
	FWD   bool
}

func main() {
	range_rover := car{}
	range_rover.Type = "Outlander"
	range_rover.Doors = 5
	range_rover.FWD = true

	pp.Println(range_rover)
}
