package main

import (
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"testing"
)

// func TestLinkParser(t *testing.T) {
// 	var tests = [][]byte{
// 		[]byte(""),
// 		[]byte(" "),
// 		[]byte("- something\n\t-child\n\t-another child\n"),
// 		[]byte("[regex101](https://regex101.com/r/mL3hA8/1)"),
//
// 		[]byte("[hugo docs content management section](https://gohugo.io/content-management/formats/)"),
// 	}
//
// 	for i := 0; i < len(tests); i++ {
// 		md := blackfriday.New()
// 		node := md.Parse(tests[i])
// 		node.Walk(Visitor)
// 	}
//
// }

func TestWithReadme(t *testing.T) {
	md := blackfriday.New()
	f, err := ioutil.ReadFile("README.md")
	if err != nil {
		t.Error(err)
	}
	node := md.Parse(f)
	node.Walk(Visitor)
}
