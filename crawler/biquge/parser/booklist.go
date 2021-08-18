package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

var bookListRe = regexp.MustCompile(`<a href="(https://www.bequ6.com/[^/]+/)">([^<]+)</a>`)

func ParseBookList(contents []byte, _ string) engine.ParseResult {
	matches := bookListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, engine.Item{})
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseBook,
		})
	}
	return result
}
