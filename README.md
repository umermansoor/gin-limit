# gin-limit
Limit the number of simultaneous HTTP connections with this Gin middleware.

## Example
```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/umermansoor/gin-limit"
)

func main() {
  r := gin.Default()
  r.Use(limit.MaxConnections(15))
  r.Run(":8080")
}
```

## Benchmarks
MBP Pro 2020, small JSON payload

### Without limit
```shell
wrk -t10 -c200 -d10s http://localhost:8080/videos

Running 10s test @ http://localhost:8080/videos
  10 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    26.75ms   26.20ms 277.71ms   58.52%
    Req/Sec     0.91k   116.67     1.25k    68.90%
  90818 requests in 10.01s, 11.95MB read
  Socket errors: connect 0, read 48, write 0, timeout 0
Requests/sec:   9070.39
Transfer/sec:      1.19MB
```

### With a limit of 15
```shell
wrk -t10 -c200 -d10s http://localhost:8080/videos

Running 10s test @ http://localhost:8080/videos
  10 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    18.94ms   20.04ms 201.40ms   85.35%
    Req/Sec     1.36k   153.70     1.93k    69.10%
  135787 requests in 10.01s, 12.46MB read
  Socket errors: connect 0, read 51, write 0, timeout 0
  Non-2xx or 3xx responses: 123310
Requests/sec:  13566.99
Transfer/sec:      1.25MB
```