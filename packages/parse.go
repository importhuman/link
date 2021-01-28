package parse

import (
	"fmt"
	"golang.org/x/net/html"
)

func TreeParser(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, val := range n.Attr {
			fmt.Println(val)
			// break
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		TreeParser(c)
	}
}
