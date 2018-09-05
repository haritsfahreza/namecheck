[![Go Report Card](https://goreportcard.com/badge/github.com/uwuh/namecheck)](https://goreportcard.com/report/github.com/uwuh/namecheck)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE)
[![Build Status](https://travis-ci.org/uwuh/namecheck.svg?branch=master)](https://travis-ci.org/uwuh/namecheck)

# namecheck
Check your name idea availability with CLI

## Installation
```
$ go get -u github.com/uwuh/namecheck/cmd/namecheck
```

## Usage
```
$ namecheck help
Usage of namecheck:
  -help
        show this message
  -name string
        name idea that you want to check

```
```
$ namecheck -name uwuh
Please wait...
V .com
V .co
V .co.id
V .id
V .info
V .io
V .me
V .name
X .net
V .org
V .us
V .xyz
V Facebook
X YouTube
V Instagram
X Twitter
V LinkedIn
V Google+
X Github

```

## License
See [LICENSE](LICENSE)