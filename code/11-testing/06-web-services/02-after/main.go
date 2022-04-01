package main

import (
	"io"
	"net/http"
)

func main() {

	// echo request body in response
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.Copy(w, r.Body)
}
