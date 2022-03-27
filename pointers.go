package main

import "fmt"
/*
  b := &a {means var b *int }
  *b means ( *&a) that is value from the addess location
*/
func main() {
	x := 0
	y = &x // means we are assigning the address of the x variable to y
	fmt.Println(*y) // means we are fetching the value from the addess of x since y holds the addess of x and has type as *int
    passByValue(x)
	fmt.Println("Pass by values")
	fmt.Println(x);
	fmt.Println("Pass by reference/ptr")
	fmt.Println("address to x" , &x)
	passByPtr(&x) // pass the address using & symbol to function
	fmt.Println("Value after changing the data by its addess")
	fmt.Println(x)

}

func passByValue(y int) {
	 x := 23

	 fmt.Println(x)
}


func passByPtr(y *int)  { // recieving the variable as pointer to the addess passing value hold 
	 *y = 45 // on the addess of y change the value to 45 i.e (*&y) 
	 
	 fmt.Println("address to passed variable",y)
	 fmt.Println("Changed value to passed variable inside the passByPtr", *y)
}