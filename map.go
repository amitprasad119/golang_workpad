package main

import "fmt"

func main() {

	// create the map 
	var ids map[int]string
   // assign the values to map
	ids = map[int]string{ 10 : "jack",
	   20: "peter",
	   30: "mack",
	   40: "rock",
	}

	fmt.Println(ids)

	// delete values from map using key

	delete(ids,30)

	fmt.Println(ids)

	// create/update another entry for map keys 

	ids[50] = "Isaac"
	ids[40] = "Milton"
	fmt.Println(ids)

	// create and initialize the map in another way as well 

	var address = make(map[string]string)
	fmt.Println(address)
	address["milton"] = "USA"
	address["issac"] = "UK"
	fmt.Println(address)

	// get the value of map 

	miltonValue := address["milton"]
	fmt.Println(miltonValue)

	// iterate over maps 

	 for key,value := range address {
		 fmt.Println("key",key , "value:",value)
	 }

}