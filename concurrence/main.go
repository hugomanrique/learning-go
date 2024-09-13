package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// canal := make(chan int)
	// canal <- 15 // send value by channel
	// valor := <- canal // receive value by channel
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}
	ch := make(chan string)
	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}

	// time.Sleep(time.Second * 2)

	alapsed := time.Since(start)
	fmt.Printf("Ejecutada Tomo %v segundos\n", alapsed.Seconds())
}

func checkAPI(url string, ch chan string) {
	if _, err := http.Get(url); err != nil {
		ch <- fmt.Sprintf("ERROR: %s no esta disponible\n", url)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s disponible\n", url)
}
