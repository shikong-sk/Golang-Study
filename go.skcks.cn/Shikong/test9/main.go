/**
指针
*/
package main

import "fmt"

func main() {
	fmt.Println("指针操作")
	a := 666
	fmt.Println("整数 a =>", a, &a)
	b := &a // 取变量 a 的地址, 将指针保存到 b 中
	fmt.Println("指针 b =>", b, &b)
	fmt.Println("将指针 b 指向的地址的值 赋值 给 c")
	c := *b
	fmt.Println("a =>", a, &a, "\nc =>", c, &c)
	fmt.Println("修改 a 的值")
	a = 999
	fmt.Println("a =>", a, &a, "\nc =>", c, &c)
	fmt.Println("修改 指针 b 指向的地址的值")
	*b = 777
	fmt.Println("a =>", a, &a, "\n指针 b =>", b, &b, "\nc =>", c, &c)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("使用 new 函数 获取类型指针")
	nInt := new(int)
	fmt.Printf("%T \t %v \t %+v\n", nInt, nInt, *nInt)
	nBool := new(bool)
	fmt.Printf("%T \t %v \t %+v\n", nBool, nBool, *nBool)
	nString := new(string)
	fmt.Printf("%T \t %v \t %q\n", nString, nString, *nString)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("使用 make 函数 分配内存")

	var m map[string]string
	m = make(map[string]string, 10)
	m["username"] = "ShiKong"
	m["age"] = "20"
	fmt.Println(m, len(m), &m)
}
