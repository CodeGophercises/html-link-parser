package main

import (
	"fmt"
	"os"

	"github.com/CodeGophercises/html-link-parser/parser"
)

func main() {
	htmlData, _ := os.ReadFile("html-samples/ex1.html")
	links, _ := parser.Parse(htmlData)
	fmt.Println(links)
}
