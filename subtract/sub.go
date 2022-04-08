package main

import "fmt"

func main() {
	result := subTwo(50, 30)
	fmt.Println(result)
}

func subTwo(a int, b int) int {
	return a - b
}
