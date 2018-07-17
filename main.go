package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "<h1>Main</h1>") })
	http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "<h1>A</h1>") })
	http.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "<h1>B</h1>") })
	http.ListenAndServe("localhost:4000", nil)
}
