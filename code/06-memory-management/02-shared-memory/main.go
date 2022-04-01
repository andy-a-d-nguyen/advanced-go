package main

import (
	"fmt"
	"sync"
)

type User struct {
	ID int
}

func (u User) String() string {
	return fmt.Sprintf("User[%v]", u.ID)
}

func main() {
	cache := make(map[int]User, 0)
	var m sync.Mutex
	const count = 1000
	var wg sync.WaitGroup
	wg.Add(count)

	for i := 0; i < 1000; i++ {
		i := i
		go func() {
			u := User{ID: i}
			m.Lock()
			cache[i%50] = User{ID: i}
			m.Unlock()
			cacheUser := cache[i%50]
			if cacheUser != u {
				fmt.Printf("Different users: %v vs %v\n", cacheUser, u)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
