package engine

import (
	"errors"
	"learngo/crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	var body []byte
	var err error
	times := 10
	for {
		log.Printf("Fetching %s\n", r.Url)
		body, err = fetcher.Fetch(r.Url)
		times--
		if times == 0 {
			err = errors.New("can not get: " + r.Url)
		}
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Printf("Fetch error %s: %v\n", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(body, r.Url)
}
