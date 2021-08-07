package parser

import (
	"learngo/crawer/engine"
	"regexp"
)

const bookRe = `<a href="(/[^/]+/[0-9]+.html)">([^<]+)</a>`

const BaseUrl = "https://www.xxbiquge.net"

func ParseBook(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(bookRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: BaseUrl + string(m[1]),
			ParseFunc: func(contents []byte) engine.ParseResult {
				return ParseChapter(contents, name)
			},
		})
	}
	return result
}
