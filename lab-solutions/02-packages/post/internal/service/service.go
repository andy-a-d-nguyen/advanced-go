package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"lab/post/internal/logic"
)

func New() http.Handler {
	return &postHandler{}
}

type postHandler struct{}

func (ph postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// split the URL path into parts. /api/posts/1 will yield ["", "api", "posts", "1"]
	pathParts := strings.Split(r.URL.Path, "/")

	// check if /api/posts
	if len(pathParts) == 3 || pathParts[3] == "" {
		switch r.Method {
		case http.MethodGet:
			ph.getAll(ctx, w, r)
		case http.MethodPost:
			ph.post(ctx, w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	} else if len(pathParts) >= 4 {
		id, err := strconv.Atoi(pathParts[3])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
			ph.get(ctx, id, w, r)
		case http.MethodDelete:
			ph.delete(ctx, id, w, r)
		case http.MethodPut:
			ph.put(ctx, id, w, r)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (ph postHandler) encodeAndReturnJSON(ctx context.Context, w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		ph.handleError(w, err, http.StatusInternalServerError)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		log.Print(err)
	}
}

func (postHandler) handleError(w http.ResponseWriter, err error, httpStatus int) {
	log.Print(err)
	w.WriteHeader(httpStatus)
}

func (ph postHandler) getAll(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	ph.encodeAndReturnJSON(ctx, w, logic.GetAll(ctx))
}
func (ph postHandler) post(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var p logic.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil || p.ID != 0 {
		ph.handleError(w, err, http.StatusBadRequest)
		return
	}
	p = logic.Add(ctx, p)
	ph.encodeAndReturnJSON(ctx, w, p)
}
func (ph postHandler) get(ctx context.Context, id int, w http.ResponseWriter, r *http.Request) {
	p, err := logic.GetByID(ctx, id)
	if err != nil {
		ph.handleError(w, err, http.StatusNotFound)
		return
	}
	ph.encodeAndReturnJSON(ctx, w, p)
}
func (ph postHandler) delete(ctx context.Context, id int, w http.ResponseWriter, r *http.Request) {
	p, err := logic.Delete(ctx, id)
	if err != nil {
		ph.handleError(w, err, http.StatusNotFound)
	}
	ph.encodeAndReturnJSON(ctx, w, p)
}
func (ph postHandler) put(ctx context.Context, id int, w http.ResponseWriter, r *http.Request) {
	var p logic.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil || p.ID != id {
		ph.handleError(w, err, http.StatusBadRequest)
		return
	}
	p, err = logic.Update(ctx, p)
	if err != nil {
		if err != nil {
			ph.handleError(w, err, http.StatusInternalServerError)
			return
		}
	}
	ph.encodeAndReturnJSON(ctx, w, p)
}
