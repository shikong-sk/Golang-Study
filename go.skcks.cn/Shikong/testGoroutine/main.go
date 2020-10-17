package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go randDelay(i)
	}
	wg.Wait()
}

func randDelay(i int) {
	defer wg.Done()
	delay := rand.Intn(1000-500) + 500
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println(i, delay)
}
