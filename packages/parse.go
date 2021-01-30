package parse

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

type Link struct {
	Href string
	Text string
}

var Links []Link

func TreeParser(n *html.Node) []Link {
	// checks if node is of type ElementNode and is <a> (initially the whole document is passed in as a node)
	if n.Type == html.ElementNode && n.Data == "a" {
		// iterates over the attributes of the node
		for _, attribute := range n.Attr {
			// check attribute type
			if attribute.Key == "href" {
				newLink := Link{}
				newLink.Href = attribute.Key
				// attribute.Val : link
				fmt.Println("newlink:", newLink)

				// iterate over node children for text
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					trimmedText := strings.TrimSpace(c.Data)
					if c.Type == html.TextNode && len(trimmedText) > 0 {
						fmt.Printf("c: %q\n", trimmedText)
						if newLink.Text == "" {
							newLink.Text = trimmedText
						} else {
							newLink.Text = newLink.Text + trimmedText
						}
						fmt.Println("newlink:", newLink)
						Links = append(Links, newLink)
						// break
					}
				}
				// break
			}
		}
	}
	// FirstChild: First child of a node
	// NextSibling: Next child of the node (at the same level)
	// Thus, this loop runs over each child of a node, calling the function on each child
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		TreeParser(c)
	}

	return Links
}

// func TextParser(n *html.Node) {
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		trimmedText := strings.TrimSpace(c.Data)
// 		if c.Type == html.TextNode && len(trimmedText) > 0 {
// 			// fmt.Printf("c: %q\n", trimmedText)
// 			newLink := Link{attribute.Val, trimmedText}
// 			Links = append(Links, newLink)
// 			break
// 		}
// 	}
// }
