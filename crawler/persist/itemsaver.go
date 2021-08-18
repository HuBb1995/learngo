package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	itemCount := 0
	go func() {
		for {
			item := <-out
			log.Printf("Got item #%d: %v\n", itemCount, item.Payload.(model.Profile).Name)
			itemCount++

			err := save(client, index, item)
			if err != nil {
				log.Printf("save item %v err: %v", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {
	indexService := client.Index().Index(index).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
