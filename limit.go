package limit

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MaxConnections(limit int) gin.HandlerFunc {
	sem := make(chan struct{}, limit)
	release := func() { <-sem }
	return func(c *gin.Context) {
		select {
		case sem <- struct{}{}: // acquire before request
			defer release() // release after request
			c.Next()
		default:
			c.AbortWithError(http.StatusServiceUnavailable,
				errors.New(fmt.Sprintf("too many connections. limit %v", limit))) // send 503 and stop the chain
		}
	}
}