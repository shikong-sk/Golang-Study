/**
函数
 */
package main

import (
	"fmt"
	"time"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	testDefer()
	fmt.Printf("\n==============================================\n\n")
	testPanic()
	fmt.Printf("\n==============================================\n\n")
	fmt.Println(time.Now())
	fmt.Printf("\n==============================================\n\n")
	fmt.Printf("testReturn()\t => \t")
	fmt.Println(testReturn())
	fmt.Printf("testReturn2()\t => \t")
	fmt.Println(testReturn2())
	fmt.Printf("\n==============================================\n\n")
	fmt.Println("函数变量")
	f := func(){
		fmt.Println("\t 函数变量 f()")
	}
	f()
}

func testDefer() {
	// 执行顺序
	// 1. x = 1, y = 2
	x := 1
	y := 2
	// 7. (1+3)返回 4, x = 10, y = 20		2. (1+2)返回 3, x = 1, y =2
	defer calc("AA", x, calc("A", x, y))
	// 3. x = 10, y = 2
	x = 10
	// 6. (10+12)返回 22, x = 10, y = 20		4. (10+2)返回 12, x = 10, y = 2
	defer calc("BB", x, calc("B", x, y))
	// 5. x = 10, y = 20
	y = 20
	fmt.Println("x =", x, "y =", y)
}

func testPanic() {
	fmt.Println("使用 defer panic 实现 try...catch")
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("错误信息:", err)
		}
	}()
	panic("test panic")
}

func testReturn()(int,int){
	a := 1
	b := 2
	return a,b
}

func testReturn2()(a,b int){
	a = 1
	b = 2
	return
}