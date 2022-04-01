package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	ph := &postHandler{}
	http.Handle("/api/posts/", ph)
	http.Handle("/api/posts", ph)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

var (
	postMutex sync.RWMutex = sync.RWMutex{}
	posts     []post       = make([]post, 0)
)

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
	ph.encodeAndReturnJSON(ctx, w, getAllPosts(ctx))
}
func (ph postHandler) post(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var p post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil || p.ID != 0 {
		ph.handleError(w, err, http.StatusBadRequest)
		return
	}
	p = addPost(ctx, p)
	ph.encodeAndReturnJSON(ctx, w, p)
}
func (ph postHandler) get(ctx context.Context, id int, w http.ResponseWriter, r *http.Request) {
	p, err := getPostByID(ctx, id)
	if err != nil {
		ph.handleError(w, err, http.StatusNotFound)
		return
	}
	ph.encodeAndReturnJSON(ctx, w, p)
}
func (ph postHandler) delete(ctx context.Context, id int, w http.ResponseWriter, r *http.Request) {
	p, err := deletePost(ctx, id)
	if err != nil {
		ph.handleError(w, err, http.StatusNotFound)
	}
	ph.encodeAndReturnJSON(ctx, w, p)
}
func (ph postHandler) put(ctx context.Context, id int, w http.ResponseWriter, r *http.Request) {
	var p post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil || p.ID != id {
		ph.handleError(w, err, http.StatusBadRequest)
		return
	}
	p, err = updatePost(ctx, p)
	if err != nil {
		if err != nil {
			ph.handleError(w, err, http.StatusInternalServerError)
			return
		}
	}
	ph.encodeAndReturnJSON(ctx, w, p)
}

type post struct {
	ID      int
	Title   string
	Content string
	Author  user
}

type user struct {
	Username string
}

func updatePost(ctx context.Context, update post) (post, error) {
	postMutex.Lock()
	defer postMutex.Unlock()
	for i, p := range posts {
		if p.ID == update.ID {
			posts[i] = update
			return update, nil
		}
	}
	return post{}, fmt.Errorf("Post with id '%v' not found", update.ID)
}

func deletePost(ctx context.Context, id int) (post, error) {
	postMutex.Lock()
	defer postMutex.Unlock()
	for i, p := range posts {
		if p.ID == id {
			posts = append(posts[:i], posts[i+1:]...)

			return p, nil
		}
	}
	return post{}, errors.New("Post not found")
}

func getPostByID(ctx context.Context, id int) (post, error) {
	postMutex.RLock()
	defer postMutex.RUnlock()
	for _, p := range posts {
		if p.ID == id {
			return p, nil
		}
	}
	return post{}, errors.New("Post not found")
}

var nextPostID = 1

func addPost(ctx context.Context, p post) post {
	postMutex.Lock()
	defer postMutex.Unlock()
	p.ID = nextPostID
	nextPostID++
	posts = append(posts, p)
	return p
}

func getAllPosts(ctx context.Context) []post {
	postMutex.RLock()
	defer postMutex.RUnlock()
	// create a copy of posts to avoid concurrent access issues
	result := make([]post, len(posts))
	copy(result, posts)
	return result
}
