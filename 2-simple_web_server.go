package main

import (
    "fmt"
    "net/http"
)

// http.ResponseWriter is the http server's response
// http.Request represents the client http request
// r.URL.Path is the path component fo the url, [1:] excludes the starting "/"

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}