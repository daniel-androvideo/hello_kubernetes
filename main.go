package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello kubernetes</h1>")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
