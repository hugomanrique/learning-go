package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return a / b, nil
}

func main() {
	// str := "123f"
	// nm, error := strconv.Atoi(str)
	// if error != nil {
	// 	fmt.Println(error)
	// 	return 
	// }
	// fmt.Println(nm)
	result, error := divide(10, 2)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(result)
}