package main

import (
	"golang.org/x/net/html"
	"os"
	"fmt"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		panic(err)
	}

	pre := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", 2 * depth, "", n.Data)
			depth++
		}
	}

	post := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", 2 * depth, "", n.Data)
		}
	}

	forEachNode(doc, pre, post)
}
