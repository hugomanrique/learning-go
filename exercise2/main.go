package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World!")

	os := runtime.GOOS
	fmt.Println("OS:", os)

	switch os {
		case "windows":
			fmt.Println("Windows")
		case "linux":
			fmt.Println("Linux")
		case "darwin":
			fmt.Println("MacOS")
		default:
			fmt.Println("Unknown")
	}
}