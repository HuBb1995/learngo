package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strings"
)

var contentRe = regexp.MustCompile(`<div id="content">(.+?)</div>`)

var idUrlRe = regexp.MustCompile(`https://www.xxbiquge.net/[^/]+/([0-9]+).html`)

func ParseChapter(contents []byte, url string, name string) engine.ParseResult {
	contentStr := extractString(contents, contentRe)
	contentStr = strings.ReplaceAll(contentStr, "&nbsp;", " ")
	contentStr = strings.ReplaceAll(contentStr, "<br />", "\n")

	chapter := model.Chapter{}
	chapter.Name = name
	chapter.Content = contentStr

	result := engine.ParseResult{}
	result.Items = append(result.Items, engine.Item{
		Id:      extractString([]byte(url), idUrlRe),
		Url:     url,
		Payload: chapter,
	})
	//file, _ := os.OpenFile("剑来.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//defer file.Close()
	//file.Write([]byte("\n\n" + name + "\n\n"))
	//file.Write([]byte(contentStr))
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
