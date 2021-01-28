package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	// file flag
	filePtr := flag.String("file", "ex2.html", "HTML file to parse")
	flag.Parse()

	// open file as bytes
	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(file)

	// parse bytefile
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(doc)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, val := range n.Attr {
				fmt.Println(val)
				// break
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
