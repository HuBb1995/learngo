package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler_distributed/config"
	"regexp"
)

var listRe = regexp.MustCompile(`<a href="(http://[^.]+.hongniang.com)" title="([^"]+)" target="_blank">[^<]+</a>`)

func ParseCityList(contents []byte, _ string) (engine.ParseResult, error) {
	matches := listRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(config.ParseCity, ParseCity),
		})
	}
	return result, nil
}
