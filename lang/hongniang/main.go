package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

var listRe = regexp.MustCompile(`<a href="(http://[^.]+.hongniang.com)" title="([^"]+)" target="_blank">[^<]+</a>`)

var cityRe = regexp.MustCompile(`<a href="(http://www.hongniang.com/user/member/id/[0-9]+)" target="_blank"> <img src="[^"]+" alt="[^"]+" title="([^"]+)">`)

var profileRe = regexp.MustCompile(`<li><span>[^：]+：</span>([^<]+)</li>`)

func main() {
	resp, err := http.Get("http://www.hongniang.com/user/member/id/11022302")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	matches := profileRe.FindAllSubmatch(body, -1)
	matches = matches[:7]
	//itemCount := 0
	for _, m := range matches {
		//itemCount++
		//fmt.Printf("#%d, Title: %s, Url: %s\n", itemCount, m[2], m[1])
		fmt.Printf("%s\n", m[1])
	}
}
