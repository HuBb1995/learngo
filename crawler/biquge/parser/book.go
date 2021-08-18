package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

var bookRe = regexp.MustCompile(`<a href="(/[^/]+/[0-9]+.html)">([^<]+)</a>`)

const BaseUrl = "https://www.xxbiquge.net"

func ParseBook(contents []byte, _ string) engine.ParseResult {
	matches := bookRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, engine.Item{})
		result.Requests = append(result.Requests, engine.Request{
			Url:       BaseUrl + string(m[1]),
			ParseFunc: BookParse(string(m[2])),
		})
	}
	return result
}

func BookParse(name string) engine.ParseFunc {
	return func(contents []byte, url string) engine.ParseResult {
		return ParseChapter(contents, url, name)
	}
}
