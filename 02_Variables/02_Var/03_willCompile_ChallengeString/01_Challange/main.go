package main

import (
	"fmt"
)

var yeti string // var yeti is of type string

func main() {
	yeti = "Abominable"
	yeti = " Snowman"
	fmt.Println(yeti)
}

// How can you get the variable yeti to print out "Abominable Snowman"?  Rather than just "Snowman"
