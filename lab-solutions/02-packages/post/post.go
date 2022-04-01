package post

import (
	"lab/post/internal/service"
	"net/http"
)

func NewHandler() http.Handler {
	return service.New()
}
