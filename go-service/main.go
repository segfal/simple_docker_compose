package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    // Update the port to 8080 (the port Python service is actually running on)
    response, err := http.Get("http://python-service:8080/api/message")
    if err != nil {
        log.Fatalf("Failed to get response: %v", err)
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalf("Failed to read response body: %v", err)
    }

    fmt.Printf("Response from Python service: %s\n", string(body))
}