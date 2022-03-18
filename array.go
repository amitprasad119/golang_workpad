package main

import "fmt"

func main() {

	names := []string{"amit", "Sumit", "suman"} // give the fix dimension

	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(cap(names))
	appendedArray := append(names, "john")
	fmt.Println(appendedArray)
}
