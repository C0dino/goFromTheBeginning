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
	// var varFromMain string // Totally Different Variable because its contained within the scope of firstFunction
	varFromMain = "Attempt to change from first function"

	return varFromMain
}

/* Change varFromMain in firstFunction to something other named variable to emphasize where its coming from */
