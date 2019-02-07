package main

import (
	// "io/ioutil"
	"fmt"
	"testing"
)

func TestLinkParser(t *testing.T) {
	// f, fileErr := ioutil.Read
	// if err != nil {
	// 	fmt.Print(err)
	// }
	var tests = []string{
		// Empty document.
		"",
		" ",
		"[regex101](https://regex101.com/r/mL3hA8/1)",

		"[hugo docs content management section](https://gohugo.io/content-management/formats/)",
	}

	for i := 0; i < len(tests); i++ {
		fmt.Println(tests[i])
		// t.Error(tests[i])
	}

}
