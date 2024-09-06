package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages map[string]int
	baseURL *url.URL
	mu *sync.Mutex
	concurrencyControl chan struct{}
	wg *sync.WaitGroup
}

func (cfg *config) crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// TODO: Remove redundant rawBaseURL and pages parameters
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}

	if baseURL.Hostname() != currentURL.Hostname() {
		fmt.Printf("skipping URL '%s' outside of base URL domain '%s'\n", currentURL, baseURL.Hostname())
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizeURL: couldn't normalize URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if _, visited := pages[normalizedCurrentURL]; visited {
		pages[normalizedCurrentURL]++
		return
	}

	pages[normalizedCurrentURL] = 1

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v\n", err)
	}

	URLs, err := getURLsFromHTML(htmlBody, rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v\n", err)
	}

	for _, URL := range URLs {
		cfg.crawlPage(rawBaseURL, URL, pages)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	return true
}