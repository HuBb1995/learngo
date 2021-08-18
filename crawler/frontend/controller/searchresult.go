package controller

import (
	"context"
	"github.com/olivere/elastic/v7"
	"learngo/crawler/engine"
	"learngo/crawler/frontend/model"
	"learngo/crawler/frontend/view"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// TODO
// fill in query string
// support search button
// rewriteQueryString
// support paging
// add start page

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

//localhost:8888/search?q=惊蛰&from=20
func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.Trim(req.FormValue("q"), " ")
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	page, err := s.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = s.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (s SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := s.client.Search("hongniang_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Query = q
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
