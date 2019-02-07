package main

import (
	"bytes"
	"fmt"
	"github.com/russross/blackfriday/v2"
	"io"
	"io/ioutil"
	"text/template"
)

type LinkRenderer struct {
	mode string
}

// empty declarations to implement the blackfriday.Renderer interface
func (r *LinkRenderer) RenderHeader(w io.Writer, ast *blackfriday.Node) {}
func (r *LinkRenderer) RenderFooter(w io.Writer, ast *blackfriday.Node) {}

type LinkData struct {
	Title string
	Dest  string
}

func dataToMap(title, dest []byte) LinkData {
	return LinkData{
		string(title),
		string(dest),
	}
}

const (
	htmlTempl = `
<li>
	{{ .Title }} <a href="{{ .Dest }}">{{ .Dest }}</a>
</li>
`
	markdownTempl = `
- {{ .Title }} [{{ .Dest }}]({{ .Dest }})
`
)

func format(mode string, title, dest []byte) []byte {
	var tmplString string
	switch mode {
	case "md":
		tmplString = markdownTempl
	case "html":
		tmplString = htmlTempl
	default:
		panic("Unrecognised template mode")
	}

	tmpl, err := template.New(mode).Parse(tmplString)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	data := dataToMap(title, dest)

	execErr := tmpl.Execute(buf, data)
	if execErr != nil {
		panic(execErr)
	}

	return buf.Bytes()
}

func (r *LinkRenderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	// only use 2nd pass
	if !entering {
		switch node.Type {
		case blackfriday.Link:
			// ignore relative links
			if node.LinkData.Destination[0] != '#' && node.LinkData.Destination[0] != '.' {
				// LinkData.Title is always empty, use FirstChild.Literal instead
				w.Write(format(r.mode, node.FirstChild.Literal, node.LinkData.Destination))
			}
		}
	}
	return blackfriday.GoToNext
}

func GenerateBibliography(mode string, input []byte) []byte {
	linkRenderer := &LinkRenderer{mode}
	output := blackfriday.Run(input, blackfriday.WithNoExtensions(), blackfriday.WithRenderer(linkRenderer))
	return output
}

func main() {
	fmt.Println(
		" ____________________ Markdown Bibliography Generator ____________________",
	)

	fmt.Println("\nParsing `README.md` links...")
	f, err := ioutil.ReadFile("README.md")
	if err != nil {
		fmt.Println(err)
	}

	output := GenerateBibliography("md", f)
	fmt.Println("Output:")
	fmt.Println(string(output))
}
