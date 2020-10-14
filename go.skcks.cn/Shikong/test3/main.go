/**
	运算符
 */
package main

import "fmt"

func main() {
	var (
		a = 1
		b = 2
	)

	fmt.Println(a, b)
	// 不使用临时变量互换变量值
	b ^= a
	a ^= b
	b ^= a
	fmt.Println(a, b, a <= b, a >= b)

	b++
	fmt.Println(a, b, a == b, a != b)

	fmt.Println(a&b, a|b, a^b, a<<b, a>>b)
}
