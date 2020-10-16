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


}
