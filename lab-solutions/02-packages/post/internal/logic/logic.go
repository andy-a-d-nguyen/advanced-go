package logic

import (
	"context"
	"errors"
	"fmt"
	"lab/user"
	"sync"
)

var (
	m     sync.RWMutex = sync.RWMutex{}
	posts []Post       = make([]Post, 0)
)

type Post struct {
	ID      int
	Title   string
	Content string
	Author  user.User
}

func Update(ctx context.Context, update Post) (Post, error) {
	m.Lock()
	defer m.Unlock()
	for i, p := range posts {
		if p.ID == update.ID {
			posts[i] = update
			return update, nil
		}
	}
	return Post{}, fmt.Errorf("Post with id '%v' not found", update.ID)
}

func Delete(ctx context.Context, id int) (Post, error) {
	m.Lock()
	defer m.Unlock()
	for i, p := range posts {
		if p.ID == id {
			posts = append(posts[:i], posts[i+1:]...)

			return p, nil
		}
	}
	return Post{}, errors.New("Post not found")
}

func GetByID(ctx context.Context, id int) (Post, error) {
	m.RLock()
	defer m.RUnlock()
	for _, p := range posts {
		if p.ID == id {
			return p, nil
		}
	}
	return Post{}, errors.New("Post not found")
}

var nextPostID = 1

func Add(ctx context.Context, p Post) Post {
	m.Lock()
	defer m.Unlock()
	p.ID = nextPostID
	nextPostID++
	posts = append(posts, p)
	return p
}

func GetAll(ctx context.Context) []Post {
	m.RLock()
	defer m.RUnlock()
	// create a copy of posts to avoid concurrent access issues
	result := make([]Post, len(posts))
	copy(result, posts)
	return result
}
