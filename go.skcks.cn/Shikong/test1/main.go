/**
	基础语法
 */
package main

import "fmt"

var (
	a int
	b int
	c int
	name string
)

// 数量级常量定义
const (
	B  = 1 << (10 * iota) // iota 是 go 语言的常量计数器，只能在常量的表达式中使用 从 0 开始计数
	KB = 1 << (10 * iota) // iota = 1
	MB = 1 << (10 * iota) // iota = 2
	GB = 1 << (10 * iota)
	_  = 1 << (10 * iota) // 使用 _(匿名变量) 跳过某些值
	PB // iota = 5 批量定义常量时,若不写赋值语句则与上一行赋值方法相同
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(name)
	fmt.Println(B, KB, MB, GB, PB)
	nameBytes := []byte("时空")
	fmt.Println(nameBytes)
	name := string(nameBytes)
	name += "666666"
	fmt.Println(name,len(name))
	fmt.Printf("name:%s\n",name)
	fmt.Printf("a:%d b:%d c:%d\n",a,b,c)
}
