package main

import "fmt"

var tarnishedShell string

func main() {

	fmt.Println(leonardo(tarnishedShell))
}

func leonardo(sword string) string {
	return donatello(sword)
}

func donatello(staff string) string {
	return raphael(staff)
}

func raphael(sai string) string {
	return michaelangelo(sai)
}

func michaelangelo(nunChucks string) string {
	return "Turtle Wax"
}
