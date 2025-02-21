package main

import (
  "fmt"
  "net/http"
  )
  
  func mesaj(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "3")
  }

func main() {
    http.HandleFunc("/", mesaj)
    http.ListenAndServe(":8080", nil)
}