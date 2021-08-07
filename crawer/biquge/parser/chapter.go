package parser

import (
	"learngo/crawer/engine"
	"os"
	"regexp"
	"strings"
	"time"
)

var introRe = regexp.MustCompile(`<div id="content">(.+?)</div>`)

//func ParseProfile(contents []byte) engine.ParseResult {
//	profile := model.Profile{}
//
//	match := introRe.FindSubmatch(contents)
//	if match != nil {
//		introS := strings.Split(string(match[1]), "|")
//		profile.City = strings.TrimSpace(introS[0])
//		profile.Age = strings.TrimSpace(introS[1])
//		profile.Education = strings.TrimSpace(introS[2])
//		profile.Marriage = strings.TrimSpace(introS[3])
//		profile.Height = strings.TrimSpace(introS[4])
//		profile.Income = strings.TrimSpace(introS[5])
//	}
//	result := engine.ParseResult{
//		Items: []interface{}{profile},
//	}
//	return result
//}

func ParseChapter(contents []byte, name string) engine.ParseResult {
	time.Sleep(time.Millisecond * 10)
	matches := introRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		item := strings.ReplaceAll(string(m[1]), "&nbsp;", " ")
		item = strings.ReplaceAll(item, "<br />", "\n")
		result.Items = append(result.Items, item)
	}
	file, _ := os.OpenFile("剑来.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	for _, item := range result.Items {
		file.Write([]byte("\n\n" + name + "\n\n"))
		itemStr := item.(string)
		file.Write([]byte(itemStr))
	}
	return result
}
