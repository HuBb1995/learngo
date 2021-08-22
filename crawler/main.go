package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/hongniang/parser"
	"learngo/crawler/persist"
	"learngo/crawler/scheduler"
	"learngo/crawler_distributed/config"
)

func main() {
	itemChan, err := persist.ItemSaver("hongniang_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	//e := engine.SimpleEngine{}
	e.Run(engine.Request{
		Url:    "http://www.hongniang.com",
		Parser: engine.NewFuncParser(config.ParseCityList, parser.ParseCityList),
	})
}
