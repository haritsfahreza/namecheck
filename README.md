[![GoDoc](https://godoc.org/github.com/uwuh/namecheck?status.svg)](https://godoc.org/github.com/uwuh/namecheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/uwuh/namecheck)](https://goreportcard.com/report/github.com/uwuh/namecheck)
[![Build Status](https://travis-ci.org/uwuh/namecheck.svg?branch=master)](https://travis-ci.org/uwuh/namecheck)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

# namecheck (`nck`)
Check your name idea availability with CLI

## Installation
```sh
$ go get -u github.com/uwuh/namecheck/cmd/nck
```

## Usage
```sh
$ nck help
Usage of nck:
  -help
        show this message
  -name string
        name idea that you want to check

```
```sh
$ nck -name uwuh

Checking uwuh
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

## License
See [LICENSE](LICENSE)