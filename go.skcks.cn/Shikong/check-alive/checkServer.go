package main

import (
	"check-alive/httpMethod"
	"check-alive/notification"
	"check-alive/printJson"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type api struct {
	address string
	alive   bool
}

func main() {
	result := make(chan string, 4)
	_ = sync.Map{}
	apiMap := new(sync.Map)
	apiMap.Store("241 Ncda 服务", api{address: "http://192.168.1.241:1080/ncda/v2/api-docs"})
	apiMap.Store("241 Ncdk 服务", api{address: "http://192.168.1.241:1089/ncdk/v2/api-docs"})
	apiMap.Store("测试 192.168.2.5 Ncda 服务", api{address: "http://192.168.2.5:8010/ncda/v2/api-docs"})
	apiMap.Store("测试 192.168.2.3 Ncdk 服务", api{address: "http://192.168.2.3:7010/ncdk/v2/api-docs"})

	processNum := 0
	apiMap.Range(func(k, v interface{}) bool {
		processNum++
		return true
	})

	for {
		wg.Add(processNum)
		apiMap.Range(func(k, v interface{}) bool {
			go process(v.(api).address, 1, result, apiMap, k.(string))
			return true
		})
		wg.Wait()
		for i := 0; i < cap(result); i++ {
			fmt.Println(<-result)
		}
		time.Sleep(time.Second)
	}
}

func process(url string, timeout int32, result chan<- string, apiMap *sync.Map, key string) {
	defer wg.Done()
	alive := CheckAlive(url, timeout)
	store, _ := apiMap.Load(key)
	if alive != store.(api).alive {
		apiMap.Store(key, api{url, alive})
		if alive {
			result <- key + " 已上线" + time.Now().Format("2006-01-02 15:04:05.000") + "\n"
			_ = notification.Notification(key+" 已上线", time.Now().Format("2006-01-02 15:04:05.000"))
		} else {
			result <- key + " 无法访问" + time.Now().Format("2006-01-02 15:04:05.000") + "\n"
			_ = notification.Notification(key+" 无法访问", time.Now().Format("2006-01-02 15:04:05.000"))
		}
	}
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
