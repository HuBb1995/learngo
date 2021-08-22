package main

import (
	"fmt"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
	"learngo/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://www.hongniang.com/user/member/id/303718",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "寻找您的人",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
