package main

import "fmt"

func main() {
	mNumber := magicNumber() // mNumber will have method as whole with type int

	fmt.Println("Got the number from magincNumber function:", mNumber()) // call the function mNumber as function by adding ()

}

func magicNumber() func() int { // `func() int` is a return type for function magicNumber()
	return func() int {
		return 42
	}
}
