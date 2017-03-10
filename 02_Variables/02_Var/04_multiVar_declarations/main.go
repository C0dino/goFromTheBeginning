package main

import (
	"fmt"
)

var (
	yeti         = "Abominable Snowman"
	donatello    = "Ninja Turtle"
	firstPlace   = 1
	skyIsFalling = false
)

func main() {

	fmt.Println("var yeti contains the string literal:", yeti)
	fmt.Printf("var yeti is of type: %T\n\n", yeti)

	fmt.Println("var donatello contains the string literal:", donatello)
	fmt.Printf("var donatello is of type: %T\n\n", donatello)

	fmt.Println("var firstPlace contains the integer:", firstPlace)
	fmt.Printf("var firstPlace is of type: %T\n\n", firstPlace)

	fmt.Println("var skyIsFalling contains the boolean value:", skyIsFalling)
	fmt.Printf("var skyIsFalling is of type: %T\n\n", skyIsFalling)
}
