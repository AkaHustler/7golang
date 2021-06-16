package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main()  {
	r := gee.New()
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Url.Path = %q\n", req.URL.Path)
	})

	r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w,"Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}