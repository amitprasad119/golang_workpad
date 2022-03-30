package main

import "fmt"

type Printer interface {
	show()
}

type Person struct {
	fname string
	lname string
}

type Vechile struct {
	model        string
	chasisNumber string
}

func (p Person) show() {
	fmt.Println("Hey there I'm ", p.fname, p.lname)
}

func (v Vechile) show() {
	fmt.Println("Vechile is  ", v.model, v.chasisNumber)
}

func main() {
	var personPrinter Printer
	var vechiclePrinter Printer
	personPrinter = Person{fname: "John", lname: "Moore"}
	personPrinter.show()
	vechiclePrinter = Vechile{model: "FX-78", chasisNumber: "FL90234324G"}
	vechiclePrinter.show()

}
