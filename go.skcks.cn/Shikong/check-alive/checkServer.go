package main

import (
	"check-alive/httpMethod"
	"check-alive/notification"
	"check-alive/printJson"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	result := make(chan string, 4)
	_ = sync.Map{}
	apiMap := new(sync.Map)
	apiMap.Store("http://192.168.1.241:1080/ncda/v2/api-docs", false)
	apiMap.Store("http://192.168.1.241:1089/ncdk/v2/api-docs",false)
	apiMap.Store("http://192.168.2.5:8010/ncda/v2/api-docs",false)
	apiMap.Store("http://192.168.2.3:7010/ncdk/v2/api-docs",false)

	processNum := 0
	apiMap.Range(func(k,v interface{})bool{
		processNum++
		return true
	})

	api := map[string]bool{
		"http://192.168.1.241:1080/ncda/v2/api-docs": false,
		"http://192.168.1.241:1089/ncdk/v2/api-docs": false,
		"http://192.168.2.5:8010/ncda/v2/api-docs":   false,
		"http://192.168.2.3:7010/ncdk/v2/api-docs":   false,
	}

	for {
		fmt.Println(time.Now())
		wg.Add(len(api))
		for k := range api {
			go process(k, 1, result, api,k)
		}
		wg.Wait()
		for i := 0; i < cap(result); i++ {
			fmt.Println(<-result)
		}
		fmt.Println()
		time.Sleep(time.Second)
	}
}

func process(url string, timeout int32, result chan<- string,api map[string]bool,key string) {
	defer wg.Done()
	alive := CheckAlive(url, timeout)
	if alive != api[key]{
		api[key] = alive
		if alive {
			_ = notification.Notification(url+"已上线", time.Now().Format("2006-01-02 15:04:05.000"))
		}
	}
	result <- url + " => " + strconv.FormatBool(alive)
}

func CheckAlive(url string, timeout int32) bool {
	_, err := httpMethod.Get(url, timeout)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func CheckAliveWithPrintJson(url string, timeout int32) bool {
	resp, err := httpMethod.Get(url, timeout)
	if err != nil {
		fmt.Println(err)
		return false
	}
	jsonMap := printJson.Json2Map(resp)
	printJson.Print(jsonMap)
	return true
}
