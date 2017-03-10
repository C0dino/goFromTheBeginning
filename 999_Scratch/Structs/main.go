package main

import (
	"fmt"
)

type car struct {
	manufacturer string
	model        string
	wheels       int
	doors        int
}

func main() {
	i := car{
		"Tesla",
		"Model S",
		4,
		4,
	}
	fmt.Println(i)
	fmt.Println("Make: ", i.manufacturer)
	fmt.Println(`Model: `, i.model)
	fmt.Println(`Wheels: `, i.wheels)
	fmt.Println(`Doors: `, i.doors)
}
