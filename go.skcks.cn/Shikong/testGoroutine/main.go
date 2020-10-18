package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// 计数器
var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go randDelay(i)
	}
	wg.Wait()

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("无缓冲通道")
	c := make(chan string)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go randDelayWithChannel(i, c)
		go getChannelData(c)
	}
	wg.Wait()
	close(c)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("有缓冲通道")
	ch := make(chan string, 5)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go randDelayWithChannel(i,ch)
		if (i+1) % cap(ch) == 0 {
			for j := 0; j < cap(ch); j++ {
				wg.Add(1)
				go getChannelData(ch)
			}
		}
	}
	wg.Wait()
}

func getChannelData(c chan string) {
	defer wg.Done()
	x := <-c
	fmt.Println(x)
}

func randDelay(i int) {
	defer wg.Done()
	delay := rand.Intn(1000-500) + 500
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println(i, delay)
}

func randDelayWithChannel(i int, c chan string) {
	defer wg.Done()
	delay := rand.Intn(1000-500) + 500
	time.Sleep(time.Duration(delay) * time.Millisecond)
	c <- strconv.Itoa(i) + "\t" + strconv.Itoa(delay)
}
