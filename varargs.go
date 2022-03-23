package main

import "fmt"

func main() {
	oneToFourTotal := addAll(1, 2, 3, 4)
	oneToTenTotal := addAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("1to4Total", oneToFourTotal)
	fmt.Println("1toTenTotal", oneToTenTotal)
}

func addAll(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}
