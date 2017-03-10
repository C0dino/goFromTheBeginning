package main

import (
	"fmt"
)

var yeti = "Abominable Snowman" // initialzing the variable word with the string literal "Hello"

var number = 1                  // Go will INFER the variable type based on the initialization

var lieDetector = true

func main() {
	fmt.Println(yeti)
	fmt.Println(number)
	fmt.Println(lieDetector)

}
