package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %w", err)
	}

	var urls []string
	var traverseErr error

	var traverseNodes func(*html.Node)
	traverseNodes = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href, err := url.Parse(attr.Val)
					if err != nil {
						traverseErr = fmt.Errorf("couldn't parse href '%s': %w", href, err)
						return
					}
					absoluteURL := baseURL.ResolveReference(href)
					urls = append(urls, absoluteURL.String())
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverseNodes(c)
		}
	}

	traverseNodes(doc)

	if traverseErr != nil {
		return nil, traverseErr
	}
		
	return urls, nil
}
