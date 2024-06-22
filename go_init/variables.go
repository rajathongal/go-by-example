package main

import "fmt"

func main() {
	// var declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	var b = 2
	fmt.Println(b)

	// You can declare multiple variables at once.
	var c, d int = 3, 4
	fmt.Println(c, d)

	// Go will infer the type of initialized variables.
	var e = true
	fmt.Println(e)

	// Variables declared without a corresponding initialization are zero-valued.
	// For example, the zero value for an int is 0.
	var f int
	fmt.Println(f)

	// The := syntax is shorthand for declaring and initializing a variable,
	// e.g. for var f string = "apple" in this case.

	// This syntax is only available inside functions.
	g := "Apple"
	fmt.Println(g)

}
