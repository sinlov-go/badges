[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov-go/badges?label=go.mod)](https://github.com/sinlov-go/badges)
[![GoDoc](https://godoc.org/github.com/sinlov-go/badges?status.png)](https://godoc.org/github.com/sinlov-go/badges)
[![goreportcard](https://goreportcard.com/badge/github.com/sinlov-go/badges)](https://goreportcard.com/report/github.com/sinlov-go/badges)
[![GitHub license](https://img.shields.io/github/license/sinlov-go/badges)](https://github.com/sinlov-go/badges)
[![codecov](https://codecov.io/gh/sinlov-go/badges/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov-go/badges)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/sinlov-go/badges)](https://github.com/sinlov-go/badges/tags)
[![GitHub release)](https://img.shields.io/github/v/release/sinlov-go/badges)](https://github.com/sinlov-go/badges/releases)

## for what

- this project used to generate badge by golang

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/sinlov-go/badges)](https://github.com/sinlov-go/badges/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## depends

in go mod project

```bash
# warning use privte git host must set∑
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "http://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q http://github.com/sinlov-go/badges.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/sinlov-go/badges
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/sinlov-go/badges | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## Features

- [X] generator support platform
  - [https://img.shields.io](https://img.shields.io)
- [X] generator different style
  - [codecov](https://codecov.io)
  - [golang](https://golang.org)
  - [npm](https://www.npmjs.com)
  - [rust](https://www.rust-lang.org)
- [x] more perfect test case coverage
- [x] more perfect benchmark case

## env

- minimum go version: go 1.17
- change `go 1.17`, `^1.17`, `1.17.13` to new go version

### libs

| lib                                 | version |
|:------------------------------------|:--------|
| https://github.com/stretchr/testify | v1.8.4  |

## usage

- see [example](https://github.com/sinlov-go/badges/tree/main/example)
- or [see pkg.go.dev](https://pkg.go.dev/github.com/sinlov-go/badges)

```go
import (
    "github.com/sinlov-go/badges/codecov_badges"
    "github.com/sinlov-go/badges/golang_badges"
    "github.com/sinlov-go/badges/npm_badges"
    "github.com/sinlov-go/badges/node_badges"
    "github.com/sinlov-go/badges/rust_badges"
    "github.com/sinlov-go/badges/shields_badges"
)
```

# dev

```bash
# It needs to be executed after the first use or update of dependencies.
$ make init dep
```

- test code

```bash
$ make test testBenchmark
```

add main.go file and run

```bash
# run at env dev use cmd/main.go
$ make dev
```

- ci to fast check

```bash
# check style at local
$ make style

# run ci at local
$ make ci
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```
