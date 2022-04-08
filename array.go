package main

import "fmt"

func main() {
	
	names := [3]string{"amit", "Sumit", "suman"} // give the fix dimension
	
	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(cap(names))
}
