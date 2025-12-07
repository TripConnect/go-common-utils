# Introduction
The system utils package for Golang  

# Installation
```sh
go get github.com/tripconnect/go-common-utils
```

# Build
Publish to Github registry
```sh
git add .
git commit -m "build: something"
git tag v<version> # ex: git tag v1.0.0
git push origin v<version> # ex: git push origin v1.0.0
```

# Unittest
Run unittest
```sh
cd go-common-utils
go test ./...
```

# Usage
## Advance search
```go
import (
	"log"

	"github.com/gocql/gocql"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/tripconnect/go-common-utils/advance_search"
)

type ChatMessageDocument struct {
	Id             gocql.UUID `json:"id"`
	ConversationId gocql.UUID `json:"conversation_id"`
	FromUserId     gocql.UUID `json:"from_user_id"`
	Content        string     `json:"content"`
	SentTime       int        `json:"sent_time"`
	CreatedAt      int        `json:"created_at"`
}


func main() {
	elasticsearchClient, _ = elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://localhost:9200")},
	})

	var musts []types.QueryVariant{
		esdsl.NewWildcardQuery("content", "something")
	}

	var esQuery types.QueryVariant = esdsl.NewBoolQuery().
		Must(musts...)

	var sort types.SortCombinationsVariant = esdsl.NewSortOptions().
		AddSortOption("field_name", esdsl.NewFieldSort(sortorder.Desc))

	searchResult, err := advance_search.NewAdvanceSearch[ChatMessageDocument]().
		Client(elasticsearchClient).
		Query(esQuery).
		Index("index_name").
		Page(0, 20).
		Sort(sort).
		Search()

	if err != nil {
		log.Printf("err: %v", err)
		return
	}

	for _, doc := range searchResult {
		log.Printf("doc: %v", doc)
	}
}


```

## Jwt
```go
...

import (
	"github.com/gofrs/uuid/v5"
    "github.com/tripconnect/go-common-utils"
)

type Claims struct {
    UserID uuid.UUID `json:"user_id"`
}

func main() {
	secret := "super-secret-123"

	userId, _ := uuid.NewV4()
	expected := Claims{UserID: userId}

	// Sign
	token, err := SignJwt(expected, secret, 5*time.Minute)
	fmt.Print(token)

	// Extract
	claims, err := ExtractJwtClaim[Claims](token, secret)
	fmt.Print(claims)
}

```

# Tips
## Local Development Setup
When developing backend services (`chat-service`, etc) together with your internal packages (`go-common-utils`, `go-proto-lib`, etc.), **never publish a new version for every tiny change**. Instead, use `replace` directives — Go will use your local folders instantly.  
Recommended `go.mod` (Windows & cross-platform safe)
```go
...

require (
	...
	github.com/tripconnect/go-common-utils v1.0.2 // real package declearation here
	...
)

replace github.com/tripconnect/go-common-utils => C:\path\trip-connect\go-common-utils // point to local package folder location
```