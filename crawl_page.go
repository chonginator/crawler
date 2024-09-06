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

func (cfg *config) crawlPage(rawCurrentURL string) {
	baseURL, err := url.Parse(cfg.baseURL.String())
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", cfg.baseURL.String(), err)
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", cfg.baseURL.String(), err)
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

	if isFirst := cfg.addPageVisit(normalizedCurrentURL); !isFirst {
		return
	}

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
		cfg.crawlPage(URL)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}