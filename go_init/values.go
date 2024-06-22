package main

import "fmt"

func main() {

	// Go has various value types including strings, integers, floats, booleans, etc.
	// Here are a few basic examples.

	// Strings, which can be added together with +.
	fmt.Println("go" + "Lang")

	// Integers and floats.
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	// Booleans, with boolean operators as you’d expect.
	fmt.Println(true && false)
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
