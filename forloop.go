package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even", i)
		} else {
			fmt.Println("Odd", i)
		}
	}

	/*
	 for loop with break
	 Create a loop and break if loop reach at 5th
	*/
	fmt.Println("Example of break with loop")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	/*
	   for loop with continue
	   Create a loop and skip if number is even and print the rest
	*/
	fmt.Println("Example of continue with loop")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
