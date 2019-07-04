# ODT

[![Build Status](https://travis-ci.org/efimovalex/odt.svg?branch=master)](https://travis-ci.org/efimovalex/odt)
[![Go Report Card](https://goreportcard.com/badge/github.com/efimovalex/odt)](https://goreportcard.com/report/github.com/efimovalex/odt) [![codecov](https://codecov.io/gh/efimovalex/odt/branch/master/graph/badge.svg)](https://codecov.io/gh/efimovalex/odt) [![GoDoc](https://godoc.org/github.com/efimovalex/odt?status.svg)](https://godoc.org/github.com/efimovalex/odt)

Golang Other Date Types

Golang Library that translates between `DATE` and `TIME` SQL types and JSON encoding.

Features: 
    - automatically converts json to mysql scannable struct for DATE and TIME types
    - can change the format read from the json
## Example:

Look in the [example folder](example/main.go)

## Change formats:

```
    odt.SetDateFormat("01/02/2006")
    odt.SetTimeFormat("22-22-22")
```
