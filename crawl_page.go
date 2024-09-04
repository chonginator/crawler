package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
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
		crawlPage(rawBaseURL, URL, pages)
	}
}