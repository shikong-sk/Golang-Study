package main

import (
	"fmt"
	"math"
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

	fmt.Println("\n==============================================")

	num2 := 0b1010 // go 1.13 版本之后可定义二进制 0b1010
	fmt.Printf("二进制：%b\n", num2)
	num8 := 012
	fmt.Printf("八进制：%o\n", num8)
	num := 10
	fmt.Printf("十进制：%d\n", num)
	num16 := 0xA
	fmt.Printf("十六进制：%X\n", num16)

	fmt.Println("\n==============================================")

	var c64 complex64
	fmt.Println("复数 complex64 实部和虚部为32位 => ", c64)
	var c128 complex128
	fmt.Println("复数 complex128 实部和虚部为64位 => ", c128)
}
