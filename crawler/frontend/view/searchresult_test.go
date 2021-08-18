package view

import (
	"learngo/crawler/engine"
	"learngo/crawler/frontend/model"
	common "learngo/crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.test.html")

	out, err := os.Create("template.test.html")
	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url: "https://www.xxbiquge.net/77_77268/336802.html",
		Id:  "336802",
		Payload: common.Chapter{
			Name:    "第一章 惊蛰",
			Content: "二月二，龙抬头。\n\n 暮色里，小镇名叫泥瓶巷的僻静地方，有位孤苦伶仃的清瘦少年，此时他正按照习俗，一手持蜡烛，一手持桃枝，照耀房梁、墙壁、木床等处，用桃枝敲敲打打，试图借此驱赶蛇蝎、蜈蚣等，嘴里念念有词，是这座小镇祖祖辈辈传下来的老话：二月二，烛照梁，桃打墙，人间蛇虫无处藏。\n\n 少年姓陈，名平安。",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
