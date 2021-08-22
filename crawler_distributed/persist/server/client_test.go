package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"learngo/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	//Start ItemSaverServer
	go ServeRpc(host, "test1")
	time.Sleep(time.Second)
	//Start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	//Call save
	result := ""
	item := engine.Item{
		Id:  "303718",
		Url: "http://www.hongniang.com/user/member/id/303718",
		Payload: model.Profile{
			Name:      "寻找您的人",
			Gender:    "女",
			City:      "浙江杭州",
			Age:       "33岁",
			Education: "高中及以下",
			Marriage:  "未婚",
			Height:    "170CM",
			Income:    "1-5万元",
		},
	}

	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s\n", result, err)
	}
}
