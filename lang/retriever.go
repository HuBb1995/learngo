package main

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.baidu.com")
}

type MockRetriever struct {
	Contents string
}

func (r MockRetriever) Get() string {
	return r.Contents
}

func main() {

}
