package parse

import (
	// "fmt"
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
				// attribute.Val : link
				newLink.Href = attribute.Val

				// fmt.Println("newlink:", newLink)

				// iterate over node children for text
				TextParser(n, newLink)
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

// iterate over each child element for text
func TextParser(n *html.Node, newLink Link) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		trimmedText := strings.TrimSpace(c.Data)
		// if it's text but not just whitespace
		if c.Type == html.TextNode && len(trimmedText) > 0 {
			// fmt.Printf("c: %q\n", trimmedText)

			// if text is not empty, adds to previous text instead of adding a new link
			if newLink.Text == "" {
				newLink.Text = trimmedText
			} else {
				newLink.Text = newLink.Text + trimmedText
			}
			// fmt.Println("newlink:", newLink)
			// break
		}
	}
	// adds to Links after all text is added
	Links = append(Links, newLink)
}
