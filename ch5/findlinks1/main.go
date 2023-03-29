package main

import (
	"fmt"
	"golang.org/x/net/html"
	// "strings"
	"os"
)

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}
func main() {
	// x := `<a href='gologolo.lo'>Link</a><a href='github.com'>github</a>`

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
