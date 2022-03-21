package main

import "fmt" 

func  main()  {
    

	var (
		num int
		num2 int
	)
	fmt.Println("Enter a number")
	fmt.Scanf("%d",&num)
   // we can directily check on the variable with expression 
	switch  {
	case num <= 10:
		  fmt.Println("Num is less than equal to 10")
		  // we can also use `fallthrough` if we want to execute the another case else by default each case will have the break
    case num <= 20: 
	      fmt.Println("Num is less than equal to 20")
	 default :
     	fmt.Println("Num is out of boundary")
	}

	fmt.Println("Enter another number(Num2):")
	fmt.Scanf("%d",&num2)
   // check the number against direct check 
	switch num2 {
	case 10 : 
	     fmt.Println("Num2 is 10")
    case 20: 
	     fmt.Println("Num2 is 20")
	case 30:
		 fmt.Println("Num2 is 30")
	 default : 
	    fmt.Println("Unknown number..")	 	 		 
	}
}