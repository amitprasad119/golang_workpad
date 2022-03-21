package main

import "fmt"

func main() {
	names := []string{"amit", "sumit", "suman"}
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10}

	entireNumbers := numbers[:]
	from4Th := numbers[3:]
	till6 := numbers[:6]
	from4to6 := numbers[3:6]

	fmt.Println(entireNumbers)
	fmt.Println(from4Th)
	fmt.Println(till6)
	fmt.Println(from4to6)
	appendedSlice := append(numbers, 67)
	fmt.Println(numbers)
	fmt.Println(appendedSlice)

	// _ is same like scala a wild card character
	for _, name := range names {
		fmt.Println(name)
	}
}
