# [Markdown link parser](#something)

At the end of every [Markdown](https://en.wikipedia.org/wiki/Markdown) article I write, I want a list of the links in the article to use as
a bibliography/list of references/useful links section.


I want to make a tool which will automagically do this for me everytime I write an article, maybe linking it with the `hugo` build command which builds my site.

Since [hugo](http://gohugo.io/) is written in [Go](http://golang.org/), I'll write the code in [Go](http://golang.org/) too.

This file aka the README file will be used to test the parser.

- [x] The first step will be to print the links in a Markdown file to stdout 
- [x] Append to the file
- [ ] Incorporate it into the hugo build process, find out how markdown yaml frontmatter like "draft: true" is processed
- [ ] Add some condition to the frontmatter, i.e. "appendlinks: bool", to toggle the behaviour

### Markdown parser: [Blackfriday](https://github.com/russross/blackfriday)

Internally, hugo also uses this package to translate markdown into html.

[hugo docs content management section](https://gohugo.io/content-management/formats/)

### Regex alternative

```regex
(?|(?<txt>(?<url>(?:ht|f)tps?://\S+(?<=\PP)))|\(([^)]+)\)\[(\g<url>)])
```

> Note: Only markdown links are included, urls like the one below are ignored

https://regex101.com/r/mL3hA8/1

