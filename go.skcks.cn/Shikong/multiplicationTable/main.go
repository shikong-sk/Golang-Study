/**
	乘法表
 */
package main

import "fmt"

func main() {
	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			fmt.Printf("%dx%d=%2d\t", x, y, x*y)
		}
		fmt.Println()
	}
}
