package advance_search

import (
	"context"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// AdvanceSearch is a fluent builder for Elasticsearch searches
type AdvanceSearch[DocType any] struct {
	client   *elasticsearch.TypedClient
	index    string
	query    *types.QueryVariant
	sort     *types.SortCombinationsVariant
	pageNum  int
	pageSize int
}

type SearchResult[DocType any] struct {
	Data          []DocType
	TotalElements int64
}

// NewAdvanceSearch creates a new search builder
func NewAdvanceSearch[DocType any]() *AdvanceSearch[DocType] {
	return &AdvanceSearch[DocType]{
		pageNum:  0,
		pageSize: 10,
	}
}

// Client sets the Elasticsearch typed client
func (s *AdvanceSearch[DocType]) Client(client *elasticsearch.TypedClient) *AdvanceSearch[DocType] {
	s.client = client
	return s
}

// Index sets the target index
func (s *AdvanceSearch[DocType]) Index(index string) *AdvanceSearch[DocType] {
	s.index = index
	return s
}

// Query sets the query (pass *types.Query, usually built with helpers)
func (s *AdvanceSearch[DocType]) Query(query types.QueryVariant) *AdvanceSearch[DocType] {
	s.query = &query
	return s
}

// Page sets page number and size
func (s *AdvanceSearch[DocType]) Page(pageNum, pageSize int) *AdvanceSearch[DocType] {
	s.pageNum = pageNum
	s.pageSize = pageSize
	return s
}

// PageSize sets only page size
func (s *AdvanceSearch[DocType]) PageSize(pageSize int) *AdvanceSearch[DocType] {
	s.pageSize = pageSize
	return s
}

// Query sets the query (pass *types.Query, usually built with helpers)
func (s *AdvanceSearch[DocType]) Sort(sort types.SortCombinationsVariant) *AdvanceSearch[DocType] {
	s.sort = &sort
	return s
}

// Search executes the search and returns the typed response
func (s *AdvanceSearch[DocType]) Search() (*SearchResult[DocType], error) {
	if s.client == nil {
		panic("Elasticsearch client is required")
	}
	if s.index == "" {
		panic("Index is required")
	}

	req := s.client.Search().
		Index(s.index)

	if s.query != nil {
		req = req.Query(*s.query)
	}

	req = req.
		Sort(*s.sort).
		From(s.pageNum * s.pageSize).
		Size(s.pageSize)

	esResp, esErr := req.Do(context.Background())

	if esErr != nil {
		return nil, esErr
	}

	docs := s.extractHitDocs(*esResp)

	searchResult := SearchResult[DocType]{
		Data:          docs,
		TotalElements: esResp.Hits.Total.Value,
	}

	return &searchResult, nil
}

// extractHitDocs extract from elaticsearch hit to generic doc type
func (s *AdvanceSearch[DocType]) extractHitDocs(resp search.Response) []DocType {
	var docs []DocType
	for _, hit := range resp.Hits.Hits {
		var doc DocType
		if err := json.Unmarshal(hit.Source_, &doc); err != nil {
			log.Printf("error parsing hit: %v", err)
			continue
		}
		docs = append(docs, doc)
	}
	return docs
}
