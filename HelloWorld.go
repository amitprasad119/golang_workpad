package main
  
import "fmt"
  

// Main function
func main() {
    a := getNumberA()
	b := getNumberB()
    fmt.Println(a+b)
	fmt.Println(addTwo(10,5))
}

func getNumberA() int {
	return 10
}

func getNumberB() int {
	return 20
}

func addTwo(x , y int) int {
	return x + y 
}