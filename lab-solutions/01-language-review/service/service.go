package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func RegisterHandlers() {
	http.HandleFunc("/", handleMessage)
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(string(b))
	w.WriteHeader(http.StatusOK)
}
