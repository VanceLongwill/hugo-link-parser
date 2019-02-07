package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLinkParser(t *testing.T) {
	var tests = [][]byte{
		[]byte(""),
		[]byte(" "),
		[]byte("- something\n\t-child\n\t-another child\n"),
		[]byte("[regex101](https://regex101.com/r/mL3hA8/1)"),

		[]byte("[hugo docs content management section](https://gohugo.io/content-management/formats/)"),
	}
	var expected = [][]byte{
		[]byte("<ul>\n</ul>"),
		[]byte("<ul>\n</ul>"),
		[]byte("<ul>\n</ul>"),
		[]byte("<ul>\n<li>regex101 - <a href=\"https://regex101.com/r/mL3hA8/1\">https://regex101.com/r/mL3hA8/1</a></li>\n</ul>"),
		[]byte("<ul>\n<li>hugo docs content management section - <a href=\"https://gohugo.io/content-management/formats/\">https://gohugo.io/content-management/formats/</a></li>\n</ul>"),
	}

	for i := 0; i < len(tests); i++ {
		html := GenerateBibiography(tests[i])
		if !bytes.Equal(html, expected[i]) {
			t.Error("Unexpected html output at position ", i)
		}
	}
}

func TestWithReadme(t *testing.T) {
	f, err := ioutil.ReadFile("README.md")
	if err != nil {
		t.Error(err)
	}
	htmlStub, htmlFileError := ioutil.ReadFile("output.html")
	if htmlFileError != nil {
		t.Error(htmlFileError)
	}

	html := GenerateBibiography(f)

	if !bytes.Equal(bytes.TrimSpace(html), bytes.TrimSpace(htmlStub)) {
		t.Error("Unexpected html output")

		fmt.Println("------------------ RECEIVED ---------------------")
		fmt.Println(string(html))
		fmt.Println("------------------ EXPECTED ---------------------")
		fmt.Println(string(htmlStub))
	}
}
