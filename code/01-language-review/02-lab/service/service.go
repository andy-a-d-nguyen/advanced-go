package service

import (
	"net/http"
)

func RegisterHandlers() {
	http.HandleFunc("/", handleMessage)
}

func handleMessage(w http.ResponseWriter, r *http.Request) {

}
