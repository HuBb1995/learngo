package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/persist"
	"learngo/crawler_distributed/rpcsupport"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must support a port")
		return
	}
	log.Fatal(ServeRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func ServeRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
