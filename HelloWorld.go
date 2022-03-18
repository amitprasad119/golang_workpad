package main
  
import "fmt"
  

// Main function
func main() {
 
    a := getNumberA()
	b := getNumberB()
    fmt.Println(a+b)
}

func getNumberA() int {
	return 10
}

func getNumberB() int {
	return 20
}