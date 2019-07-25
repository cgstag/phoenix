# ⚡ Phoenix Golang Web Application

[![release](https://img.shields.io/badge/release%20-v0.1-0077b3.svg?style=flat-square)](https://gitlab.com/itau-im/basic/releases)

## Intro

Basic is aimed at providing a project structure and samples of every base component an app should have :
 * Router
 * Logger
 * HTTP Handlers
 * Nested Structs
 * Configuration from Yaml file and ENV
 * Tests
 * JSON encode and decode (in progress)
 * Database calls (in progress)
 * HTTP External calls (in progress)
 * HTTP Internal calls (in progress)
 
 ## Quick Start
 
 `Before anything make sure you have Go 1.11+ environment installed`.  
Link to docs : [https://golang.org/dl/](https://golang.org/dl/).  

```bash
$ git clone git@gitlab.com:deroo/phoenix.git
$ cd phoenix
$ go build
$ go run phoenix
```

### With docker

```bash
$ docker build -t deroo/phoenix .
$ docker run deroo/phoenix
```

## Test it

### Unit-tests

```bash
go test ./...
```

### Postman

Collection in progress.

`/` return Hello World
`/health` return health check
`/v1/account/random` return a generated random account