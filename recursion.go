package main

import "fmt"

func main() {

	f := factorial(5)
	fmt.Println(f)
}

func factorial(n int) int {
	if n == 1 {
		return n
	}
	return n * factorial(n-1)
}
