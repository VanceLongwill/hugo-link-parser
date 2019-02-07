# [Markdown link parser](#something)

I started writing articles for my blog 
At the end of every article I write, I want a list of the links in the article
The article will be written in [Markdown](https://en.wikipedia.org/wiki/Markdown)

I want to make a tool which will automagically do this for me everytime I write an article, maybe linking it with the `hugo` build command which builds my site.

Since [hugo](http://gohugo.io/) is written in [Go](http://golang.org/), I'll write the code in [Go](http://golang.org/) too.

This file aka the README file will be used to test the parser

[] - The first step will be to print the links in a Markdown file to stdout 

[] - then they can be appended to the file

[] - then I can figure out how to incorporate it into the build process, find out how "draft: bool" works

[] - then they I can add some condition to the frontmatter, i.e. "appendlinks: bool", to toggle the behaviour

### A markdown parser: [Blackfriday](https://github.com/russross/blackfriday)

Internally, hugo also uses this package to translate markdown into html.

[hugo docs content management section](https://gohugo.io/content-management/formats/)

### A regex alternative

```regex
(?|(?<txt>(?<url>(?:ht|f)tps?://\S+(?<=\PP)))|\(([^)]+)\)\[(\g<url>)])
```

https://regex101.com/r/mL3hA8/1

