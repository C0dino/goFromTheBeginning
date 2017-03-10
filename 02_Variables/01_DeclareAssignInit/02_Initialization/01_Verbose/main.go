package main

import (
	"fmt"
)

var yeti string = "Abominable Snowman" // initialzing the variable word with the string literal "Abominable Snowman"
var number int = 1                     // Go will INFER the variable type based on the initialization
var lieDetector bool = true

func main() {
	fmt.Println(yeti)
	fmt.Println(number)
	fmt.Println(lieDetector)

}
