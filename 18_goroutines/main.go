package main

import (
	"fmt"
	"net/http"
	"sync"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	defer resp.Body.Close()

	cType := resp.Header.Get("Content-Type")
	fmt.Printf("%s -> %s\n", url, cType)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://apy.github.com",
		"https://httpbin.org/xml",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}
