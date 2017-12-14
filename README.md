[golang-stats-api-handler](https://github.com/fukata/golang-stats-api-handler) available for frame work

## Install
```
go get github.com/y-ogura/stats-handler
```

## Example
```
import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/y-ogura/stats-handler"
)

func main() {
	e := echo.New()
	e.GET("/stats", stats_handler.EchoStatsHandler)
	e.Start(":8000")
}
```

## Response
```
$ curl -i localhost:8000/stats
HTTP/1.1 200 OK
Content-Type: text/plain; charset=UTF-8
Date: Thu, 14 Dec 2017 13:43:20 GMT
Content-Length: 520

{
  "time": 1513259000865514903,
  "go_version": "go1.8.3",
  "go_os": "darwin",
  "go_arch": "amd64",
  "cpu_num": 4,
  "goroutine_num": 4,
  "gomaxprocs": 4,
  "cgo_call_num": 1,
  "memory_alloc": 383784,
  "memory_total_alloc": 383784,
  "memory_sys": 3084288,
  "memory_lookups": 15,
  "memory_mallocs": 5071,
  "memory_frees": 140,
  "memory_stack": 327680,
  "heap_alloc": 383784,
  "heap_sys": 1769472,
  "heap_idle": 950272,
  "heap_inuse": 819200,
  "heap_released": 917504,
  "heap_objects": 4931,
  "gc_next": 4473924,
  "gc_last": 0,
  "gc_num": 0,
  "gc_per_second": 0,
  "gc_pause_per_second": 0,
  "gc_pause": []
}
```
