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
				linkPointer := &newLink
				// attribute.Val : link
				linkPointer.Href = attribute.Val

				// iterate over node children for text
				TextParser(n, linkPointer)
				// fmt.Println("pointer outside parser:", linkPointer)

				// fixes spacing issues by:
				// 1. strings.Fields splits text wherever multiple whitespaces are present into a slice
				// 2. strings.Join joins the slice objects with a single space
				linkPointer.Text = strings.Join(strings.Fields(linkPointer.Text), " ")
				// add to Links after all the text has been added
				Links = append(Links, *linkPointer)
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
func TextParser(n *html.Node, linkPointer *Link) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// not necessarily text, could be tag if it's an element node
		trimmedText := strings.TrimSpace(c.Data)
		// if it's text but not just whitespace
		if c.Type == html.TextNode && len(trimmedText) > 0 {
			// fmt.Printf("c: %q\n", trimmedText)

			// if text is not empty, adds to previous text instead of adding a new link
			if linkPointer.Text == "" {
				linkPointer.Text = c.Data
			} else {
				linkPointer.Text = linkPointer.Text + c.Data
			}
			// fmt.Println("pointer in parser:", linkPointer)
		}
		if c.Type == html.ElementNode && c.FirstChild != nil {
			// linkPointer.Text = linkPointer.Text + " "
			TextParser(c, linkPointer)
		}
	}
	// If I append here, then same link is repeated due to recursion, so adding outside the function instead
	// Links = append(Links, *linkPointer)
}
