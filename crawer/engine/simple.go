package engine

import (
	"errors"
	"learngo/crawer/fetcher"
	"log"
)

type SimpleEngine struct{}

func (*SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)
		//for _, item := range parseResult.Items {
		//	log.Printf("Got item %v\n", item)
		//}
	}

}

func worker(r Request) (ParseResult, error) {
	var body []byte
	var err error
	times := 10
	for {
		log.Printf("Fetching %s\n", r.Url)
		body, err = fetcher.Fetch(r.Url)
		times--
		if times == 0 {
			return ParseResult{}, errors.New("can not get: " + r.Url)
		}
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Printf("Fetch error %s: %v\n", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}
