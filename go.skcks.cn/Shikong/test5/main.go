/**
数组
*/
package main

import "fmt"

func main() {
	// 使用默认的值初始化
	var arr [5]int
	fmt.Println(arr)

	// 使用指定的初始值完成初始化
	var arr2 = [5]int{1, 2, 3}
	fmt.Println(arr2)

	// 让编译器自动推断数组长度
	arr3 := [...]int{1, 2, 3}
	fmt.Println(arr3)

	// 使用指定索引值的方式初始化数组
	arr4 := [...]int{1: 2, 4: 5}
	fmt.Println(arr4)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("遍历数组的方式：")
	fmt.Println("1.	for 循环")
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	fmt.Println("2.	for range")
	for index, value := range arr3 {
		fmt.Println("索引：", index, "值：", value)
	}

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("二维数组")
	arr5 := [3][2]string{
		{"北京", "北京"},
		{"广东", "广州"},
		{"浙江", "杭州"},
	}
	fmt.Println(arr5)

	fmt.Println("多维数组")
	arr6 := [4][3][2]int{{{1,2},{3,4},{5,6}},{{7,8},{9,10}}}
	fmt.Println(arr6)
}
