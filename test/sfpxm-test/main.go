package main

import (
	"fmt"
	"net/http"
	"sfpxm"
)

func main() {
	r := sfpxm.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		// %q
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":9999")
}
