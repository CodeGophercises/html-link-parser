package main

import (
	"fmt"
	"os"

	"github.com/CodeGophercises/html-link-parser/parser"
)

// Sample app to test parser
func main() {
	htmlData, _ := os.ReadFile("html-samples/ex1.html")
	links, _ := parser.Parse(htmlData)
	for _, link := range links {
		fmt.Println(link.Href)
	}
}
