# gin-limit
Limit the number of simultaneous HTTP connections with this Gin middleware.

## Example
```go
package main

import (
  "github.com/umansoor/gin-limit"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(limit.MaxConnections(100))
  r.Run(":8080")
}
```

