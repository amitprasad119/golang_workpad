package main

import "fmt"



type country struct {
	countryname string
	countrycode string
}

type address struct {
	street string
	countryInfo country
}

func main() {
  india := country{countryname:"india",countrycode:"+91"}
  var us country

  us.countryname = "United States of America"
  us.countrycode = "+1"
  
  bostonCity := address{street:"7th cross",countryInfo: country{countryname:"USA",countrycode:"+91"}}
  fmt.Println(india)
  fmt.Println(us)
  fmt.Printf("%+v",us)
  fmt.Println()
  fmt.Printf("%+v",bostonCity)
}