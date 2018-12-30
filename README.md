[![GoDoc](https://godoc.org/github.com/haritsfahreza/namecheck?status.svg)](https://godoc.org/github.com/haritsfahreza/namecheck)
[![Build Status](https://travis-ci.org/haritsfahreza/namecheck.svg?branch=master)](https://travis-ci.org/haritsfahreza/namecheck)
[![codecov](https://codecov.io/gh/haritsfahreza/namecheck/branch/master/graph/badge.svg)](https://codecov.io/gh/haritsfahreza/namecheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/haritsfahreza/namecheck)](https://goreportcard.com/report/github.com/haritsfahreza/namecheck)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

# namecheck (`nck`)
Check your name idea availability with CLI

## Installation
```sh
$ go get -u github.com/haritsfahreza/namecheck/cmd/nck
```

## Usage
```sh
$ nck help
Usage of nck:
  -code string
        specific channel code that you want to check
  -help
        show this message
  -list
        show available channel code list
  -name string
        your name idea that you want to check

```
```sh
$ nck -name haritsfahreza

Checking haritsfahreza
Please wait...

Status
V Available
X Not Available
? Unknown (e.g. Timeout)

V .org
V .io
V .co.id
V .name
V .co
V .us
V .xyz
V .info
X Twitter
V Google+
X Github
X .net
V Facebook
V .me
V .id
X YouTube
V .com
V Instagram

Duration: 2.099445s

```

## Run in Docker
```sh
$ docker build -t namecheck .
$ docker run -it namecheck -name foo
```

## License
See [LICENSE](LICENSE)
