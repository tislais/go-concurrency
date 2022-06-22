package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		if b, ok := queryCache(id); ok {
			fmt.Println("\033[0;35m from cache \033[0m")
			fmt.Println(b)
			continue
		}

		if b, ok := queryDatabase(id); ok {
			fmt.Println("\033[0;32m from database \033[0m")
			fmt.Println(b)
			continue
		}

		fmt.Printf("\033[0;36m book not found with id: %v \033[0m\n", id)
		time.Sleep(150 * time.Millisecond)
	}
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	// simulates the time it takes for the database to perform action
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}

	return Book{}, false
}
