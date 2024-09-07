package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(argsWithoutProg) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	rawBaseURL := argsWithoutProg[0]

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing base URL '%s': %v", baseURL, err)
	}

	const maxConcurrency = 10

	cfg := config{
		make(map[string]int),
		baseURL,
		&sync.Mutex{},
		make(chan struct{}, maxConcurrency),
		&sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go func() {
		cfg.concurrencyControl <-struct{}{}
		defer func() {
			<-cfg.concurrencyControl
			cfg.wg.Done()
		}()
		cfg.crawlPage(cfg.baseURL.String())
	}()
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}