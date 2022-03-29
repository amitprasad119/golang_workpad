package main

import "fmt"

type Person struct {
	fname string
	lname string
}

func main() {
	james := Person{fname: "James", lname: "Bond"}
	james.speak()
}

func (p Person) speak() {
	fmt.Println("Hi my name is ", p.fname, p.lname)
}
