package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	// everytime we iterate, we're going to generate a goroutine
	// to query our cache and a goroutine to query our database
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		// pass id in to prevent shared memory problems
		go func(id int) {
			if b, ok := queryCache(id); ok {
				fmt.Println("\033[0;35m from cache \033[0m")
				fmt.Println(b)
			}
		}(id)

		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("\033[0;32m from database \033[0m")
				fmt.Println(b)
			}
		}(id)

		// Ff we comment out the sleep here, we get no result when running the app
		// because the main function doesnt have anything to pause itself.
		// So while it is generating 20 goroutines, there isnt enough time for
		// those goroutines to return.
		// This is because program execution completes when the main function exits.
		// So we run through the loop and generate 20 goroutines and then we're done.
		// The sleep call gives time for the goroutine to finish.
		// Although we may not get all 20 because the last two goroutines aren't
		// going to have time to execute.
		// time.Sleep(150 * time.Millisecond)
	}

	// guarantees all 20 goroutines will have time to execute, assuming they take less than 2s
	time.Sleep(2 * time.Second)
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
