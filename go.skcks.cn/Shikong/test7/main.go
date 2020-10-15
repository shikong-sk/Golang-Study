package main

import "fmt"

func main() {
	m := make(map[string]int, 8)
	fmt.Println("map 类型定义 map[KeyType]ValueType 默认初始值为 nil 需使用 make 分配内存")
	fmt.Println(m)
	m["age"] = 20
	m["score"] = 100
	fmt.Println(m)

	fmt.Println("map 类型也支持在 声明的时候填充元素")
	m2 := map[string]string{
		"username": "ShiKong",
		"age":      "20",
	}
	fmt.Println(m2)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("判断键是否存在")
	val, exists := m2["password"]
	if exists {
		fmt.Println(val)
	} else {
		fmt.Println("password 不存在")
	}

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("map 使用 for...range 遍历")
	for k, v := range m2 {
		fmt.Println(k, "=>", v)
	}

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("使用 delete 函数 根据键名 删除指定的键值对")
	delete(m2, "age")
	fmt.Println(m2)

}
