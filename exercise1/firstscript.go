package main

import (
	"fmt"
	"math"
	"strconv"

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

	var (
		defaultInt int
		defaultUint uint
		defaultFloat float64
		defaultString string
		defaultBool bool
		defaultComplex complex128
		defaultRune rune
	
	)

	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultString, defaultBool, defaultComplex, defaultRune)

	//CAST TYPES DATA
	var integer16 int16 = 10
	var integer32 int32 = 100
	fmt.Println(int32(integer16)+ integer32)

	var s string = "100"
	vlr,_ := strconv.Atoi(s) // Pasar de string a int
	fmt.Println(vlr + vlr)
	var n int = 42
	vlr2 := strconv.Itoa(n) // Pasar de int a string
	fmt.Println(vlr2 + vlr2)
}