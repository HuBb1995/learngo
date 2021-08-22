package worker

import (
	"errors"
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/hongniang/parser"
	"learngo/crawler_distributed/config"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type CrawlService struct{}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Requests []Request
	Items    []engine.Item
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeSerializeRequest(r Request) (engine.Request, error) {
	engineParser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, nil
	}
	return engine.Request{
		Url:    r.Url,
		Parser: engineParser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeSerializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request; %v\n", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(config.ParseCityList, parser.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(config.ParseCity, parser.ParseCity), nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid args: %v", p.Args)
		}
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}
