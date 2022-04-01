package main

import (
	"lab/post"
	"log"
	"net/http"
)

func main() {
	ph := post.NewHandler()
	http.Handle("/api/posts/", ph)
	http.Handle("/api/posts", ph)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
