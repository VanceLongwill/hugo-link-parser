package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestWithReadme(t *testing.T) {
	f, err := ioutil.ReadFile("README.md")
	if err != nil {
		t.Error(err)
	}

	htmlStub, htmlFileError := ioutil.ReadFile("output.html")
	if htmlFileError != nil {
		t.Error(htmlFileError)
	}

	mdStub, mdFileError := ioutil.ReadFile("output.md")
	if mdFileError != nil {
		t.Error(mdFileError)
	}

	html := GenerateBibliography("html", f)
	md := GenerateBibliography("md", f)

	testEqual := func(a, b []byte) {
		if !bytes.Equal(bytes.TrimSpace(a), bytes.TrimSpace(b)) {
			t.Error("Unexpected output")

			fmt.Println("------------------ RECEIVED ---------------------")
			fmt.Println(string(a))
			fmt.Println("------------------ EXPECTED ---------------------")
			fmt.Println(string(b))
		}
	}

	testEqual(html, htmlStub)
	testEqual(md, mdStub)
}
