package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

//handler echo URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(w, "URL.path = %q\n", req.URL.Path)
}

//handler echo URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request)  {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}