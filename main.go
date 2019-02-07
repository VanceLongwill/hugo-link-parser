package main

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"io"
	"io/ioutil"
)

type LinkInfo struct {
	Title       string
	Destination string
}
type LinkRenderer struct {
	// The output should be written to the supplied writer w. If your
	// implementation has no header to write, supply an empty implementation.
	// RenderHeader(w io.Writer, ast *blackfriday.Node)

	// RenderFooter is a symmetric counterpart of RenderHeader.
	// RenderFooter(w io.Writer, ast *blackfriday.Node)
}

func Visitor(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	// only use 2nd pass
	if !entering {
		switch node.Type {
		case blackfriday.Link:
			// ignore relative links
			if node.LinkData.Destination[0] != '#' && node.LinkData.Destination[0] != '.' {
				// fmt.Println(string(node.LinkData.Title))
				// LinkData.Title is always empty, use FirstChild.Literal instead
				fmt.Println(string(node.FirstChild.Literal), string(node.LinkData.Destination))
			}
		}
	}
	return blackfriday.GoToNext
}

func (r *LinkRenderer) RenderHeader(w io.Writer, ast *blackfriday.Node) {}
func (r *LinkRenderer) RenderFooter(w io.Writer, ast *blackfriday.Node) {}
func (r *LinkRenderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	// only use 2nd pass
	if !entering {
		switch node.Type {
		case blackfriday.Link:
			// ignore relative links
			if node.LinkData.Destination[0] != '#' && node.LinkData.Destination[0] != '.' {
				// fmt.Println(string(node.LinkData.Title))
				// LinkData.Title is always empty, use FirstChild.Literal instead
				// fmt.Println(string(node.FirstChild.Literal), string(node.LinkData.Destination))
				w.Write([]byte("\n"))
				w.Write(node.FirstChild.Literal)
				w.Write([]byte("\n"))
				w.Write(node.LinkData.Destination)
				w.Write([]byte("\n"))
			}
		}
	}
	return blackfriday.GoToNext
}

func GetLinks(input []byte) []LinkInfo {

	linkRenderer := &LinkRenderer{}
	output := blackfriday.Run(input, blackfriday.WithNoExtensions(), blackfriday.WithRenderer(linkRenderer))
	fmt.Println(string(output))

	return []LinkInfo{LinkInfo{"a", "b"}}
}

func main() {
	fmt.Println(
		" ____________________ Markdown Link Parser ____________________",
	)

	f, err := ioutil.ReadFile("README.md")
	if err != nil {
		fmt.Println(err)
	}

	a := GetLinks(f)
	for i := 0; i < len(a); i++ {
		// fmt.Println(a[i])
	}
}
