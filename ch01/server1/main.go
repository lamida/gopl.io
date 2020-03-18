package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	f := func(w http.ResponseWriter, r *http.Request) { // cozy way to define anonimous function
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	}
	http.HandleFunc("/", f) // actually even we don't need temporary variable at all
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
