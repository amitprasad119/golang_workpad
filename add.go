package main

import "fmt"

// Main function
func main() {
	a := getNumberA()
	b := getNumberB()
	fmt.Println(a + b)
	fmt.Println(addTwo(10, 5))

	var (
		num1 int
		num2 int
	)
	fmt.Println("Enter num1")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter num2")
	fmt.Scanf("%d", &num2)

	output := num1 - num2

	fmt.Println("result:", output)

}

func getNumberA() int {
	return 10
}

func getNumberB() int {
	return 20
}

func addTwo(x, y int) int {
	return x + y
}
