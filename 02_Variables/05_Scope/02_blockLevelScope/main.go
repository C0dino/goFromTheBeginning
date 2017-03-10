package main

import (
	"fmt"
)

func main() {
	varFromMain := "Variable from main"
	fmt.Println(varFromMain)
	fmt.Println(firstFunction())
}

func firstFunction() string {
	varFromFirstFunction := "Variable from first function"
	return varFromFirstFunction
}
