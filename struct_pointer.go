package main

import "fmt" 

/** 
Pointer 
-------
Go is pass by value language by default 
However if you want to pass the address of the variable following are the steps 

use 
1. &<variable>
2. create variable with *ptr 
3. access (*ptr).field = <"value">

**/

type person struct {
	name string
	id int 
	city string
}

func main() {
   jack := person {name:"jack", id: 120, city:"bangalore"}

   fmt.Println(jack)
   jack.updateName("jacky")
   jack.print()

   // after passed by pointer 
   jackPtr := &jack // assign the address reference 
   jackPtr.updateNameByPtr("jacky")
   jackPtr.printByPtr()
   
}

func (p person) updateName(name string) {
	p.name = name
}

func (p person) print() {
	fmt.Println("After passed by values")
	fmt.Println(p)
}

func (personPtr *person) updateNameByPtr(name string) {
	(*personPtr).name = name // access the ptr with * i.e (*ptr).property
}

func (personPtr *person) printByPtr() {
	fmt.Println("after passed by reference")
	fmt.Println(personPtr)
}