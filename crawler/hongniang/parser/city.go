package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://www.hongniang.com/user/[^"]+)" target="_blank"> <img src="[^"]+" alt="[^"]+" title="[^"]+"> <p class="[^"]+">([^<]+)</p>`)

var nextPageRe = regexp.MustCompile(`<a style="[^"]+" class="next" href="([^"]+)">[^<]+</a>`)

func ParseCity(contents []byte, url string) (engine.ParseResult, error) {
	matches := cityRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseProfile,
		})
	}

	matches = nextPageRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//result.Items = append(result.Items, engine.Item{})
		result.Requests = append(result.Requests, engine.Request{
			Url:       url + string(m[1]),
			ParseFunc: ParseCity,
		})
	}

	return result, nil
}

func ParseProfileWithName(name string) engine.ParseFunc {
	return func(contents []byte, url string) (engine.ParseResult, error) {
		return ParseProfile(contents, url)
	}
}
