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

	"github.com/tripconnect/go-common-utils/advance_search"
)

func main() {
	searchResult, err := advance_search.NewAdvanceSearch[models.ChatMessageDocument]().
		Client(common.ElasticsearchClient).
		Query(esQuery).
		Index(consts.ChatMessageIndex).
		PageSize(int(req.GetLimit())).
		Sort(sort).
		Search()
	
	log.Printf("%v - %v", err, searchResult)
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