package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Node 6 (Go)")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8081", nil)
}
