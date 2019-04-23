# Shorty

## Description
Shorty is an application client to shortener your long url on shortener URL provider. Currently only support [kutt](https://kutt.it) provider.

## Diagram Flow
```
<user> - <shorty> - <shortenerProvider>
```
Basically it just client interact with shortener provider.

## Installation
```
go get -u github.com/ardikabs/shorty
```

## Usage
### `func GetListURL() ([]URLDefinition, error)`
```golang

package main

import (
    "fmt"
    "log"

    "github.com/ardikabs/shorty/kutt"
)

func main(){

    api := kutt.API{
        BaseURL: &url.URL{
            Scheme: "https",
            Host: "kutt.it",
        },
        APIToken: "your-api-token-here",
    }

    listURLs, err := api.GetListURL()
    
    if err != nill {
        log.Fatalln(err)
    }

    fmt.Println(listURLs)

}
```
### `func SubmitURL(longURL, customURL, password string, reuse bool) (URLDefinition, error)`
```golang

package main

import (
    "fmt"
    "log"

    "github.com/ardikabs/shorty/kutt"
)

func main(){

    api := kutt.API{
        BaseURL: &url.URL{
            Scheme: "https",
            Host: "kutt.it",
        },
        APIToken: "your-api-token-here",
    }

    resultURL, err := api.SubmitURL(
        "https://example.com",
        "my-custom-example",
        nil,
        false,
    )
    
    if err != nill {
        log.Fatalln(err)
    }

    fmt.Println(resultURL)
}
```
### `func DeleteURL(targetURL string) error`
```golang

package main

import (
    "fmt"
    "log"

    "github.com/ardikabs/shorty/kutt"
)

func main(){

    api := kutt.API{
        BaseURL: &url.URL{
            Scheme: "https",
            Host: "kutt.it",
        },
        APIToken: "your-api-token-here",
    }

    err := api.DeleteURL("https://kutt.it/my-example-room")
    
    if err != nill {
        log.Fatalln(err)
    }
}
```