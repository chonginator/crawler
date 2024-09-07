package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", cfg.baseURL.String(), err)
		return
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		fmt.Printf("skipping URL '%s' outside of base URL domain '%s'\n", currentURL, cfg.baseURL.Hostname())
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

	URLs, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v\n", err)
	}

	for _, URL := range URLs {
		cfg.wg.Add(1)
		go func(URL string) {
			cfg.concurrencyControl <-struct{}{}
			defer func() {
				<-cfg.concurrencyControl
				cfg.wg.Done()
			}()
			cfg.crawlPage(URL)
		}(URL)
	}
}
