package main

import "fmt"

func main() {
	name, age := getPersonDetails()

	fmt.Printf("%s is %d years old\n", name, age)
	
}

func getPersonDetails() (string, int) {
	return "John", 29
}
