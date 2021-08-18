package engine

type ParseFunc func(contents []byte, url string) (ParseResult, error)

type Request struct {
	Url       string
	ParseFunc ParseFunc
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

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
