package main

import (
	"fmt"
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
	f, err := ioutil.ReadFile("README.md")
	if err != nil {
		t.Error(err)
	}

	a := GetLinks(f)
	for i := 0; i < len(a); i++ {
		fmt.Println(a)
	}
}
