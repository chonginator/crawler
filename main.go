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

	cfg := config{
		make(map[string]int),
		baseURL,
		&sync.Mutex{},
		make(chan struct{}),
		&sync.WaitGroup{},
	}

	cfg.crawlPage(cfg.baseURL.String())

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}