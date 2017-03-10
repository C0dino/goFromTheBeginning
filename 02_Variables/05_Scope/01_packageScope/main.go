package main

import (
	"fmt"
)

var packageLevel = "Usable in all Functions"

func main() {
	fmt.Println(packageLevel) //packageLevel is available here
	fmt.Println(firstFunction())
}

func firstFunction() string {
	return packageLevel //packageLevel is also available here
}
