package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"packages/packages"
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

	// parse the html in the reader (not necessarily .html file) and return the parse tree (this parse tree is not in a readable format)
	// doc is of type *html.Node
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(doc)

	result := parse.TreeParser(doc)
	fmt.Println(result)
}
