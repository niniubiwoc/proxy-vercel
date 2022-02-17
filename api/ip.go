package handler

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("http://myip.ipip.net/")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Fprintf(w, string(body))
}
