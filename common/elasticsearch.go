package common

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
	"github.com/tripconnect/go-common-utils/helper"
)

var ElasticsearchClient *elasticsearch.TypedClient

func init() {
	host, hostErr := helper.ReadConfig[string]("database.elasticsearch.host")

	if hostErr != nil {
		log.Fatalf("failed to load elasticsearch config")
	}

	var err error
	ElasticsearchClient, err = elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:9200", host)},
	})

	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}
}

func GetResponseDocs[T any](resp *search.Response) []T {
	var docs []T
	for _, hit := range resp.Hits.Hits {
		var doc T
		if err := json.Unmarshal(hit.Source_, &doc); err != nil {
			log.Printf("error parsing hit: %v", err)
			continue
		}
		docs = append(docs, doc)
	}
	return docs
}
