/*
切片
*/
package main

import "fmt"

func main() {
	var a []int
	fmt.Println(a)

	b := []int{1, 2, 3, 4, 5}
	fmt.Println(b)

	arr := [...]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}
	fmt.Println("数组之间比较", arr, arr2, "=>", arr == arr2)
	fmt.Println("切片为引用值 不能直接比较 只能与 nil 比较")
	fmt.Println("a == nil", a == nil)
	fmt.Println("b == nil", b == nil)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("切片表达式[low:high] b[1:3] 取 切片 b", b, "下标[1,3) 的值", b[1:3])

	fmt.Println("切片表达式[low:high:max] b[1:3:5] 取 切片 b", b, "下标[1,3) 的值 容量为 (max - low)", b[1:3:5], cap(b[1:3:5]))

	fmt.Printf("\n==============================================\n\n")

	c := make([]int, 5, 10)
	fmt.Println("使用 make 函数 构造 元素数量:", len(c), "容量:", cap(c), " 的切片", c)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("使用 len(s) 判空 而不能用 s == nil")
	var s1 []int
	s2 := []int{}
	s3 := make([]int, 0)
	fmt.Println(len(s1), cap(s1), s1 == nil)
	fmt.Println(len(s2), cap(s2), s2 == nil)
	fmt.Println(len(s3), cap(s3), s3 == nil)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("直接将切片赋值给另一个切片，那么他们将共用同一个底层数组")
	s4 := make([]int, 5)
	s5 := s4
	s5[0] = 666
	fmt.Println("s4", s4)
	fmt.Println("s5", s5)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("遍历切片的方式与数组相同：")
	fmt.Println("1.	for 循环")
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	fmt.Println("2.	for range")
	for index, value := range b {
		fmt.Println("索引：", index, "值：", value)
	}

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("使用 append 函数 动态添加元素：")
	fmt.Println(b)
	fmt.Println(append(b, 4))
	fmt.Println(append(b, 5, 6, 7, 8))

	fmt.Println("使用 copy 函数复制切片")
	d := make([]int, 5, 5)
	fmt.Println(b, d)
	copy(d, b)
	fmt.Println("修改通过 copy 赋值的切片不会修改到原来的切片")
	d[0] = 666
	fmt.Println(b, d)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("从切片中删除元素")
	fmt.Println(d, d[:2], d[2:4], d[4:])
	d = append(d[:2], d[4:]...)
	fmt.Println(d)
}
