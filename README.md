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
```sh
cd go-common-utils
go test ./...
```

# Usage
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