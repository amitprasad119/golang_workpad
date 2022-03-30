package main

import "fmt"

func main() {

	// Type 1:
	func() {
		fmt.Println("I am anonymous function")
	}() // () meaning we are calling the function explicitly and function would not have any name

	// Type 2: assign anonymous function to a variable
	fn := func() {
		fmt.Println("Hello I am assigned function to a variable fn")
	}

	fn() // call fn with () to invove the anonymous function

	// Type 3: parameterised anonymous function
	magicNumber := func(number int) {
		fmt.Println("I got the magic number now", number)
	}

	magicNumber(42) // passing 42 as magic number

}
