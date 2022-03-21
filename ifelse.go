package main

import "fmt"

func main() {
	var num int

	fmt.Println("Enter the number")
	fmt.Scanf("%d", &num)

	if num%2 == 0 {
		fmt.Printf("Number %d is divisible by 2", num)
	} else if num%3 == 0 {
		fmt.Printf("Number %d is divisible by 3", num)
	} else {
		fmt.Println("Number %d is unknown", num)
	}

}
