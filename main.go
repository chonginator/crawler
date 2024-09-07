package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println("no website provided")
		return
	}
	if len(argsWithoutProg) > 3 {
		fmt.Println("too many arguments provided")
		return
	}
	rawBaseURL := argsWithoutProg[0]
	
	maxConcurrency := 3
	if len(argsWithoutProg) == 2 {
		var err error
		maxConcurrency, err = strconv.Atoi(argsWithoutProg[1])
		if err != nil {
			fmt.Printf("unable to parse max concurrency command-line arg")
			os.Exit(1)
		}
	}

	maxPages := 100
	if len(argsWithoutProg) == 3 {
		var err error
		maxPages, err = strconv.Atoi(argsWithoutProg[2])
		if err != nil {
			fmt.Printf("unable to parse max pages command-line arg")
			os.Exit(1)
		}
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURL)
}