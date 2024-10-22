package cmd

import (
    "fmt"
    "net/http"
)

func Execute() {

    resp, err := http.Get("http://example.com/")
    if err != nil {
        fmt.Printf("%s", err)
    } else {
        fmt.Printf("%v", resp)
    }
}
