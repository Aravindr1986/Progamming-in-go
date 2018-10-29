package main

import "fmt"

/*
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.
	go f("goroutine")
	f("direct")

	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.

	// You can also start a goroutine for an anonymous
	// function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in
	// separate goroutines now, so execution falls through
	// to here. This `Scanln` requires we press a key
	// before the program exits.
	fmt.Scanln()
	fmt.Println("done")
}*/

func printText(mesg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(mesg)
	}
}
func main() {

	go printText("From Routine1")

	printText("from main")
	go printText("From Routine33")
	fmt.Scanln()
	fmt.Println("done")

}
