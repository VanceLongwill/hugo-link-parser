package main

import (
	"fmt"
	"github.com/russross/blackfriday/v2"
	"io"
	"io/ioutil"
)

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

type LinkRenderer struct{}

// empty declarations to implement the blackfriday.Renderer interface
func (r *LinkRenderer) RenderHeader(w io.Writer, ast *blackfriday.Node) {
	w.Write([]byte("<ul>"))
}
func (r *LinkRenderer) RenderFooter(w io.Writer, ast *blackfriday.Node) {
	w.Write([]byte("\n</ul>"))
}

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
				w.Write([]byte("\n<li>"))
				w.Write(node.FirstChild.Literal)
				w.Write([]byte(" - <a href=\""))
				w.Write(node.LinkData.Destination)
				w.Write([]byte("\">"))
				w.Write(node.LinkData.Destination)
				w.Write([]byte("</a>"))
				w.Write([]byte("</li>"))
			}
		}
	}
	return blackfriday.GoToNext
}

func GenerateBibiography(input []byte) []byte {
	linkRenderer := &LinkRenderer{}
	output := blackfriday.Run(input, blackfriday.WithNoExtensions(), blackfriday.WithRenderer(linkRenderer))
	return output
}

func main() {
	fmt.Println(
		" ____________________ Markdown Link Extractor ____________________",
	)

	f, err := ioutil.ReadFile("README.md")
	if err != nil {
		fmt.Println(err)
	}

	output := GenerateBibiography(f)
	fmt.Println(string(output))
}
