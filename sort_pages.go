package main

import "sort"

type pageCount struct{
	normalizedURL string
	count int
}

func sortPages(pages map[string]int) []pageCount {
	pagesSlice := []pageCount{}
	for normalizedURL, count := range pages {
		pagesSlice = append(pagesSlice, pageCount{
			normalizedURL: normalizedURL,
			count: count,
		})
	}
	sort.Slice(pagesSlice, func(i, j int) bool {
		if pagesSlice[i].count == pagesSlice[j].count {
			return pagesSlice[i].normalizedURL < pagesSlice[j].normalizedURL
		}
		return pagesSlice[i].count > pagesSlice[j].count
	})
	return pagesSlice
}