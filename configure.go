package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct{
	pages map[string]int
	baseURL *url.URL
	mu *sync.Mutex
	concurrencyControl chan struct{}
	wg *sync.WaitGroup
	maxPages int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) shouldCrawl() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages) < cfg.maxPages
}

func configure(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL '%s': %v", rawBaseURL, err)
	}

	cfg := config{
		make(map[string]int),
		baseURL,
		&sync.Mutex{},
		make(chan struct{}, maxConcurrency),
		&sync.WaitGroup{},
		maxPages,
	}

	return &cfg, nil
}