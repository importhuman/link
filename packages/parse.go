package parse

import (
	"fmt"
	"golang.org/x/net/html"
)

func TreeParser(n *html.Node) {
	// checks if node is of type ElementNode and is <a> (initially the whole document is passed in as a node)
	if n.Type == html.ElementNode && n.Data == "a" {
		// iterates over the attributes of the node
		for _, val := range n.Attr {
			fmt.Println(val)
			// break
		}
	}
	// FirstChild: First child of a node
	// NextSibling: Next child of the node (at the same level)
	// Thus, this loop runs over each child of a node, calling the function on each child
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		TreeParser(c)
	}
}
