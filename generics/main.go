package main

import (
	"fmt"
)

func PrintList(list ...any){
	for _, value := range list {
		fmt.Println(value)
	}
}

func Sums[T Numeros](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

type integer int
type Numeros interface {
	int| ~float64 | ~float32	
}

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

type Product[T int | string] struct {
	Id int
	Desc string
	Price float32
}

func main(){
	PrintList("Hugo","Manrique")
	PrintList(22,3,4,5,6)
	total := Sums(4.5,7,1)

	strings := []string{"a","b","c"}
	fmt.Println(Includes(strings,"a"))
	fmt.Println(total)
}