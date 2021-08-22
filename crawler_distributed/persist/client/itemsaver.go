package client

import (
	"learngo/crawler/engine"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	itemCount := 0
	go func() {
		for {
			item := <-out
			log.Printf("Got item #%d: %v\n", itemCount, item)
			itemCount++
			//Call rpc to save item
			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("save item %v err: %v", item, err)
			}
		}
	}()
	return out, nil
}
