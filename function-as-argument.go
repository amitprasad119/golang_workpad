package main

import "fmt"

func greet(name string) string {
	return "Hello " + name + "!"
}

// g is of type function which takes a argument as string and return as string and second argument is name
func greetAndPrint(g func(name string) string, name string) {
	fmt.Println(g(name)) // callig finally g(name) which will internally call greet
}

func main() {
	greetAndPrint(greet, "Amit") // sending greet function as an argument as signature is satifying it
}
