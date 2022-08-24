# bgin
[![Go Report Card](https://goreportcard.com/badge/github.com/foursking/bgin)](https://goreportcard.com/report/github.com/foursking/bgin)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/foursking/bgin.svg?branch=master)](https://travis-ci.org/foursking/bgin) 
[![Foundation](https://img.shields.io/badge/Golang-Foundation-green.svg)](http://golangfoundation.org) 
[![GoDoc](https://pkg.go.dev/badge/github.com/foursking/bgin?status.svg)](https://pkg.go.dev/github.com/foursking/bgin?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/foursking/bgin/-/badge.svg)](https://sourcegraph.com/github.com/foursking/bgin?badge)
[![Release](https://img.shields.io/github/release/foursking/bgin.svg?style=flat-square)](https://github.com/foursking/bgin/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/foursking/bgin)](https://www.tickgit.com/browse?repo=github.com/foursking/bgin)
[![goproxy.cn](https://goproxy.cn/stats/github.com/foursking/bgin/badges/download-count.svg)](https://goproxy.cn)

bgin is a API framework written in Go (Golang). 
An MVCS, Restful, and version control framework based on the [Gin](https://github.com/gin-gonic/gin) framework.
If you need performance and good productivity, you will love bgin.

## Installation
To install bgin package, you need to install Go and set your Go workspace first.

1. The first need Go installed (version 1.11+ is required), then you can use the below Go command to install bgin.
```bash
$ go get -u github.com/foursking/bgin
```

2. Import it in your code:
```go 
import "github.com/foursking/bgin" 
```

## Quick start

```go
package main

import (
	"fmt"
	"github.com/foursking/bgin"
	"github.com/foursking/bgin/consts"
	"github.com/foursking/bgin/handler"
	"github.com/foursking/bgin/helper/config"
	"github.com/foursking/bgin/middleware"
	"github.com/foursking/bgin/option"
	"github.com/gin-gonic/gin"
	"demo/route"
)

func main() {
	gin.DisableConsoleColor()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {}

	if err := config.ReadFile(); err != nil {
		panic(fmt.Errorf("read config file error: %s", err))
	}

	app := bgin.New(
		option.WithMode(consts.DevMode),
		option.WithSignSecretKey("#$%1234"),
	)

	app.Run(
		handler.WithNoRoute(),
		handler.WithHealth(),
		middleware.WithZapRecovery(),
		middleware.WithZapLogger(),
		handler.WithExpVar(),
		handler.WithPrometheus(),
		handler.WithSwagger(),
		handler.WithPProf(),
		middleware.WithRestCheck(route.Restful),
	)
}
```
run main.go and visit http://0.0.0.0:9010/health (for windows "http://localhost:8080/health") on browser
```bash
$ go run main.go
```

## bgin v1. stable
- MVCS four-tier architecture support
- Restful interface style support
- API version control and permission custom configuration
- PProf middleware support
- Prometheus middleware support
- Swagger api docs middleware support
- Graceful server shutdown and reload

## Build with jsoniter
bgin uses encoding/json as default json package but you can change to [jsoniter](https://github.com/json-iterator/go) by build from other tags.
```bash
$ go build -tags=jsoniter .
```

## License
Released under the [MIT License](https://github.com/foursking/bgin/blob/master/LICENSE)



# bgin
