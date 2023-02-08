// Get content type of sites
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func siteSerial(urls []string) {
	for _, url := range urls {
		returnType(url)
	}
}

// implementation of a concurrent version
func sitesConcurrent(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		// add 1 for each go routine
		wg.Add(1)
		//spin a go-routine
		go func(url string) {
			returnType(url)
			// signal done
			wg.Done()
		}(url)
	}
	// wait for all go-routines
	wg.Wait()
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}

	start := time.Now()
	siteSerial(urls)
	fmt.Println(time.Since(start))

	start = time.Now()
	sitesConcurrent(urls)
	fmt.Println(time.Since(start))
}
