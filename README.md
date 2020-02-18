# varorequest [![Go Report Card](https://goreportcard.com/badge/github.com/varolib/varorequest)](https://goreportcard.com/report/github.com/varolib/varorequest) [![GoDoc](https://godoc.org/github.com/varolib/varorequest?status.svg)](https://godoc.org/github.com/varolib/varorequest)

## Motivations

This is a practice library to build request for a bank. The library is suppose to be minimalistic to avoid any type of bugs and unexpected behavior. 

## Golang version

`varorequest` is currently compatible with golang version from 1.12+.

## Request Builder [![GoDoc](https://godoc.org/github.com/varolib/varorequest#ReqBuilder)
```go
package main
  
import (
    "fmt"
    request "github.com/varolib/varorequest"
    "io/ioutil"
)

func main() {
    rb := request.RequestBuilder{}
    resp, err := rb.SetURL("https://google.com").Do()
    if err != nil {
        panic(err)
    }
    f, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(f))
}
```
