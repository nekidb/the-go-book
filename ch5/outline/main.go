package main

import (
	"fmt"
	"golang.org/x/net/html"
	// "strings"
	"os"
)

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Println(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func main() {
	// siteContent := "<div>Dabudi</div>"

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	outline(nil, doc)
}
