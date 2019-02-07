package main

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
)

// (?|(?<txt>(?<url>(?:ht|f)tps?://\S+(?<=\PP)))|\(([^)]+)\)\[(\g<url>)])

func Visitor(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	// fmt.Println(node.Type)
	switch node.Type {
	case blackfriday.Link:
		fmt.Println(string(node.LinkData.Destination), string(node.LinkData.Title))
	}
	return 0
}

func main() {
	fmt.Println(
		" ____________________ Markdown Link Parser ____________________",
	)
}
