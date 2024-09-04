package main

import (
	"fmt"
	"math"

	"rsc.io/quote"
)


func main(){
	fmt.Println("holal mundo")
	fmt.Println(quote.Hello())
	firstName, lastName, age := "Hugo","Manrique", 25
	fmt.Println(firstName, lastName, age)

	var integer int = 10
	var float float64 = 10.5
	fmt.Println(integer, float)
	fmt.Println(math.MaxInt64, math.MaxInt64)
}