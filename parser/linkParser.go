package parser

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	href string
	text string
}

func formatLink(link Link) {
	fmt.Printf("Href: %s\n", link.href)
	fmt.Printf("Text: %s\n\n", strings.TrimSpace(link.text))
}

// Result Set
var links []Link

// Assuming html is UTF-8 encoded html input
func Parse(htmlData []byte) ([]Link, error) {
	//  The idea is that we will parse the HTML data into syntax tree. We can use a parser for this or write our own. Will just use one
	// for now. Then traverse that tree with DFS and process any target NODES.

	rootNode, err := html.Parse(bytes.NewReader(htmlData))
	if err != nil {
		return nil, errors.New("error while parsing the html file into syntax tree")
	}

	// At this point, we have a syntax tree for the HTML doc. We just to walk it in a way efficient for our purpose.
	// Lets do DFS.
	var dfs func(*html.Node, bool, *[]byte)
	dfs = func(n *html.Node, insideAnchorTag bool, childTextContainer *[]byte) {
		// Process current node

		isAnchorTag := n.Type == html.NodeType(3) && n.Data == "a"
		isTextNode := n.Type == html.NodeType(1)

		if isAnchorTag {
			// have a container only for yourself
			*childTextContainer = make([]byte, 0)
		}

		// Am I inside <a> ?
		if insideAnchorTag {
			if isTextNode {
				// I should really put my text stuff into my parents container
				*childTextContainer = append(*childTextContainer, []byte(n.Data)...)
			}
		}

		// Go through each child
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			dfs(c, isAnchorTag || insideAnchorTag, childTextContainer)
		}

		// if current node is <a> ,when done with all children, you see your conatiner. For now, you print it.
		if isAnchorTag {
			link := Link{}
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link.href = attr.Val
					break
				}
			}
			link.text = string(*childTextContainer)
			// Append link to resultSet
			links = append(links, link)
		}
	}
	var childTextContainer []byte
	dfs(rootNode, false, &childTextContainer)
	return links, nil
}
