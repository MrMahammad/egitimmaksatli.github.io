package main

import "net/http"
  
  func anasayfa(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./anasayfa.html")
  }

func main() {
    http.HandleFunc("/", anasayfa)
    http.ListenAndServe(":8080", nil)
}
