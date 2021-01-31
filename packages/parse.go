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
			fmt.Println("newlink at start of if loop:", newLink)
			// fmt.Printf("c: %q\n", trimmedText)

			// if text is not empty, adds to previous text instead of adding a new link
			if newLink.Text == "" {
				newLink.Text = trimmedText
			} else {
				newLink.Text = newLink.Text + trimmedText
			}
			// fmt.Println("newlink at end of if loop:", newLink)
			// fmt.Println("Links in if loop:", Links)
			// fmt.Println("----------------")
			// break
		}
		if c.Type == html.ElementNode && c.FirstChild != nil {
			// fmt.Println("else loop:", c.Data)
			TextParser(c, newLink)
			// for d := c.FirstChild; d != nil; d = d.NextSibling {
			// 	trimmedText = strings.TrimSpace(d.Data)
			// 	if d.Type == html.TextNode && len(trimmedText) > 0 {
			// 		fmt.Println("child data:", d.Data)
			// 	}
			// }
		}
	}
	// adds to Links after all text is added
	Links = append(Links, newLink)
	// fmt.Println("Links outside loop:", Links)
	// fmt.Println("---------------------------------------------")
}
