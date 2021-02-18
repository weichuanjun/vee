package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(wi http.HandlerFunc, req *http.Request) {
	fmt.Fprint(w, "URL path is %q/n", req.URL.Path)
}
