package engine

import "learngo/crawler_distributed/config"

type ParserFunc func(contents []byte, url string) (ParseResult, error)

type Parser interface {
	Parse(contents []byte, url string) (ParseResult, error)
	Serialize() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Id      string
	Url     string
	Payload interface{}
}

type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) (ParseResult, error) {
	return ParseResult{}, nil
}

func (NilParser) Serialize() (name string, args interface{}) {
	return config.NilParser, nil
}

type FuncParser struct {
	name   string
	parser ParserFunc
}

func NewFuncParser(name string, parser ParserFunc) *FuncParser {
	return &FuncParser{
		name:   name,
		parser: parser,
	}
}

func (f *FuncParser) Parse(contents []byte, url string) (ParseResult, error) {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}
