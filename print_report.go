package main

import "fmt"

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)
	sortedPages := sortPages(pages)
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.count, page.normalizedURL)
	}
}