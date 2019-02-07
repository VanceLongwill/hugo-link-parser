package main

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
)

func Visitor(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	// only use 2nd pass
	if !entering {
		switch node.Type {
		case blackfriday.Link:
			// ignore relative links
			if node.LinkData.Destination[0] != '#' {
				// fmt.Println(string(node.LinkData.Title))
				// LinkData.Title is always empty, use FirstChild.Literal instead
				fmt.Println(string(node.FirstChild.Literal), string(node.LinkData.Destination))
			}
		}
	}
	return blackfriday.GoToNext
}

func main() {
	fmt.Println(
		" ____________________ Markdown Link Parser ____________________",
	)
}
