/**
变量类型
*/
package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

func main() {
	var uInt8 uint8 // 相当于 C语言中的 byte
	fmt.Println("无符号 8位 整型 0 ~ 255 => ", uInt8, math.MaxUint8)
	var uInt16 uint16 // 相当于 C语言中的 short
	fmt.Println("无符号 16位 整型 0 ~ 65535 => ", uInt16, math.MaxUint16)
	var uInt32 uint32
	fmt.Println("无符号 32位 整型 0 ~ 4294967295 => ", uInt32, math.MaxUint32)
	var uInt64 uint64 // 相当于 C语言中的 long
	fmt.Println("无符号 64位 整型 0 ~ 18446744073709551615 => ", uInt64, uint64(math.MaxUint64))
	fmt.Println("有符号 8位 整型 -128 ~ 127 => ", math.MinInt8, math.MaxInt8)
	fmt.Println("有符号 16位 整型 -32768 ~ 32767 => ", math.MinInt16, math.MaxInt16)
	fmt.Println("有符号 32位 整型 -2147483648 ~ 2147483647 => ", math.MinInt32, math.MaxInt32)
	fmt.Println("有符号 64位 整型 -9223372036854775808 ~ 9223372036854775807 => ", math.MinInt64, math.MaxInt64)

	var uInt uint
	fmt.Println("特殊类型 uint => 在32位系统上是 uint32, 64位操作系统上是 uint64", uInt)
	var iInt int
	fmt.Println("特殊类型 int => 在32位系统上是 int32, 64位操作系统上是 int64", iInt)
	var uIntPtr uintptr
	fmt.Println("特殊类型 uintptr => 无符号整型, 用于存放一个指针", uIntPtr)

	fmt.Printf("\n==============================================\n\n")

	num2 := 0b1010 // go 1.13 版本之后可定义二进制 0b1010
	fmt.Printf("二进制：%b\n", num2)
	num8 := 012
	fmt.Printf("八进制：%o\n", num8)
	num := 10
	fmt.Printf("十进制：%d\n", num)
	num16 := 0xA
	fmt.Printf("十六进制：%X\n", num16)

	fmt.Printf("\n==============================================\n\n")

	var f32 float32
	fmt.Println("浮点型 32位 =>", f32, math.MaxFloat32)
	var f64 float64
	fmt.Println("浮点型 64位 =>", f64, math.MaxFloat64)

	fmt.Printf("\n==============================================\n\n")

	var c64 complex64
	fmt.Println("复数 complex64 实部和虚部为32位 => ", c64)
	var c128 complex128
	fmt.Println("复数 complex128 实部和虚部为64位 => ", c128)

	fmt.Printf("\n==============================================\n\n")

	var b bool
	fmt.Println("布尔型 只有 true 和 false (go 中不允许将整型强制转换为布尔型) => ", b)

	var str string
	// `...` 为多行字符串 在多行字符串中 转义字符无效 所有字符原样输出
	str = `
	'\r' => 回车符（返回行首)
	'\n' => 换行符（直接跳到下一行的同列位置） 
	'\t' => 制表符
	'\"' => 双引号 
	'\\' => 反斜杠`
	fmt.Println("字符串 string =>", str)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("强制类型转换 int -> float64 =>", math.Sqrt(float64(3*3+4*4)))

	fmt.Printf("\n==============================================\n\n")

	str1 := "123一二三"
	fmt.Println(
		"统计字符串长度的方法：("+str1+")",
		"\n len => ", len(str1),
		"\n bytes.Count => ", bytes.Count([]byte(str1), nil)-1,
		"\n strings.Count => ", strings.Count(str1, "")-1,
		"\n utf8.RuneCountInString => ", utf8.RuneCountInString(str1),
	)
	strArr := strings.Split(str1, "")
	fmt.Println("分割字符串：\n", strArr)
	fmt.Println("拼接字符串：\n", strings.Join(strArr, ""))
	fmt.Println("字符串是否包含子串：\n", strings.Contains(str1, "6"))
	fmt.Println("字符串是否包含前缀：\n", strings.HasPrefix(str1, "123"))
	fmt.Println("字符串是否包含后缀：\n", strings.HasSuffix(str1, "一二三"))

	fmt.Println("替换字符串：\n", strArr)
	str2 := []rune(str1)
	fmt.Println(str2[3:6], string(str2))
	copy(str2[3:6], []rune("666"))
	fmt.Println(str2, string(str2))
}
