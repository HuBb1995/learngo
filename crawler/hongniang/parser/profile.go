package parser

import (
	"errors"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strings"
)

var profileRe = regexp.MustCompile(`<li><span>[^：]+：</span>([^<]+)</li>`)

var idRe = regexp.MustCompile(`<div class="loveid">Loveid:[0-9]+</div>`)

var nameRe = regexp.MustCompile(`<div class="name nickname" >([^"]+)<`)

var moreRe = regexp.MustCompile(`<a target="_blank" href="(/user/[^"]+)">`)

func ParseProfile(contents []byte, url string) (engine.ParseResult, error) {
	matches := profileRe.FindAllSubmatch(contents, -1)

	if len(matches) == 0 {
		return engine.ParseResult{}, errors.New("parse error: can not match anything")
	}
	profile := model.Profile{}
	profile.Name = extractString(contents, nameRe)
	profile.Age = string(matches[0][1])
	profile.Marriage = string(matches[1][1])
	profile.Height = string(matches[2][1])
	profile.Education = string(matches[3][1])
	profile.Income = string(matches[4][1])
	profile.City = string(matches[5][1])
	profile.Gender = string(matches[6][1])

	result := engine.ParseResult{}
	result.Items = append(result.Items, engine.Item{
		Id:      extractString(contents, idRe),
		Url:     url,
		Payload: profile,
	})
	//file, _ := os.OpenFile("剑来.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//defer file.Close()
	//file.Write([]byte("\n\n" + name + "\n\n"))
	//file.Write([]byte(contentStr))

	matches = moreRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		url = "http://www.hongniang.com" + string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:       url,
			ParseFunc: ParseProfile,
		})
	}

	return result, nil
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return strings.Trim(string(match[1]), " ")
	}
	return ""
}
