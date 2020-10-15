/**
流程控制
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		a = 1
		b = 2
	)

	a ^= b

	if a > b {
		fmt.Println("a > b", a, b)
	} else if a == b {
		fmt.Println("a == b", a)
	} else {
		fmt.Println("b > a", b, a)
	}

	var arr []int
	for i := 0; i < 10; i++ {
		arr = append(arr, 9-i)
	}
	fmt.Println(arr)

	// 相当于其他语言的 while
	for {
		fmt.Println(time.Now())
		time.Sleep(time.Duration(500) * time.Millisecond)
		fmt.Println(time.Now())
		break
	}

	// 遍历对象
	for k, v := range arr {
		fmt.Println("key:", k, "value:", v)
	}

	fmt.Printf("\n==============================================\n\n")

	// 生成随机数种子
	rand.Seed(time.Now().UnixNano())
	fmt.Println("随机数种子 =>",time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10-1) + 1)
	}
}