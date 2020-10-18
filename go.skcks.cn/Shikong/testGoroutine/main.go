package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// 计数器
var wg sync.WaitGroup

func main() {
	fmt.Println("可用逻辑核心数 =>", runtime.GOMAXPROCS(0))

	fmt.Printf("\n==============================================\n\n")

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
		go randDelayWithChannel(i, ch)
		if (i+1)%cap(ch) == 0 {
			for j := 0; j < cap(ch); j++ {
				wg.Add(1)
				go getChannelData(ch)
			}
		}
	}
	wg.Wait()
	close(ch)

	fmt.Printf("\n==============================================\n\n")

	testSelect()

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("并发 资源抢夺")
	var x int64
	add := func(x *int64) {
		for i := 0; i < 5000; i++ {
			*x = *x + 1
		}
		wg.Done()
	}
	wg.Add(2)
	go add(&x)
	go add(&x)
	wg.Wait()
	fmt.Println(x)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("原子操作版 资源抢夺 供了底层的原子级内存操作,比锁更轻量")
	var atomicInt int64
	add = func(aInt *int64) {
		for i := 0; i < 5000; i++ {
			atomic.AddInt64(aInt,1)
		}
		wg.Done()
	}
	wg.Add(2)
	go add(&atomicInt)
	go add(&atomicInt)
	wg.Wait()
	fmt.Println(atomicInt)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("并发 互斥锁版 资源抢夺")
	x = 0
	var lock sync.Mutex
	add = func(x *int64) {
		for i := 0; i < 5000; i++ {
			lock.Lock()
			*x = *x + 1
			lock.Unlock()
		}
		wg.Done()
	}
	wg.Add(2)
	go add(&x)
	go add(&x)
	wg.Wait()
	fmt.Println(x)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("读写 互斥锁 适用于读多写少的场景")
	rwLock := new(sync.RWMutex)
	read := func() {
		// 加读锁
		rwLock.RLock()
		time.Sleep(time.Millisecond)
		rwLock.RUnlock()
		wg.Done()
	}
	write := func() {
		// 加写锁
		rwLock.Lock()
		time.Sleep(time.Millisecond * 10) // 模拟写入延迟
		rwLock.Unlock()
		wg.Done()
	}
	startTime := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	endTime := time.Now()
	fmt.Println(endTime.Sub(startTime))

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("sync.Map 线程安全的 map 类型")
	syncMap := new(sync.Map)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			syncMap.Store(i, i+1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	syncMap.Range(func(k, v interface{}) bool {
		fmt.Println(k, "=>", v)
		return true
	})
}

func getChannelData(c <-chan string) { // <- chan string 只读 单向通道
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

func randDelayWithChannel(i int, c chan<- string) { // c chan <- string 只写 单向通道
	defer wg.Done()
	delay := rand.Intn(1000-500) + 500
	time.Sleep(time.Duration(delay) * time.Millisecond)
	c <- strconv.Itoa(i) + "\t" + strconv.Itoa(delay)
}

func testSelect() {
	fmt.Println("select 语句 \n" +
		"可处理一个或多个channel的发送/接收操作\n" +
		"如果多个case同时满足，select会随机选择一个\n" +
		"对于没有case的select{}会一直等待，可用于阻塞main函数")

	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(i, "接收 :", x)
		case ch <- i:
		}
	}
}
