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

var channelData chan string

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go randDelay(i)
	}
	wg.Wait()

	//
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go getChannelData()
	//	go randDelayWithChannel(i)
	//}
	//wg.Wait()
}

func getChannelData(){
	data := make([]string, 10)
	x := <-channelData
	fmt.Println(x)
	fmt.Println(data)
}

func randDelay(i int) {
	defer wg.Done()
	delay := rand.Intn(1000-500) + 500
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println(i, delay)
}

func randDelayWithChannel(i int) {
	defer wg.Done()
	delay := rand.Intn(1000-500) + 500
	time.Sleep(time.Duration(delay) * time.Millisecond)
	channelData <- strconv.Itoa(i) + "\t" + strconv.Itoa(delay)
}
