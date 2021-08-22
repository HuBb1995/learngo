package main

import (
	"flag"
	"learngo/crawler/engine"
	"learngo/crawler/hongniang/parser"
	"learngo/crawler/scheduler"
	"learngo/crawler_distributed/config"
	itemsaver "learngo/crawler_distributed/persist/client"
	"learngo/crawler_distributed/rpcsupport"
	worker "learngo/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma seperated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor, err := worker.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	//e := engine.SimpleEngine{}
	e.Run(engine.Request{
		Url:    "http://www.hongniang.com",
		Parser: engine.NewFuncParser(config.ParseCityList, parser.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
