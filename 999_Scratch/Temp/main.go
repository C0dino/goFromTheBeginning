package main

import (
	"fmt"
)

func main() {

	z := add(.2, .4)

	fmt.Println(z)

	fmt.Printf("%.1f", z)
	if z == 0.6 {
		fmt.Println("\n True: z is equal to 0.6")
	} else {
		fmt.Println("\nFalse: z is not equal to 0.6")
	}

	if z == 0.6000000000000001 {
		fmt.Println("z is equal to 0.6000000000000001")
	} else {
		fmt.Println("z is NOT equal to 0.6000000000000001")
	}

}

func add(x, y float64) float64 {
	return x + y
}
