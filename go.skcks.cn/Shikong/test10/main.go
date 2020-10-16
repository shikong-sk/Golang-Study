/**
结构体
*/
package main

import "fmt"

type myInt int
type s string

// 结构体定义
type person struct {
	name        string
	age, height int
}

func (p *person) SetName(newName string){
	p.name = newName
}

func (p person) SetName2(newName string) person {
	p.name = newName
	return p
}

func main() {
	var a myInt
	var s s
	fmt.Println("类型别名")
	fmt.Printf("%T \t %v \n", a, a)
	fmt.Printf("%T \t\t %q \n", s, s)

	fmt.Printf("\n==============================================\n\n")

	fmt.Println("结构体")
	var p1 person
	p1.name = "时光"
	p1.age = 20
	p1.height = 169
	p2 := person{"时空", 20, 168}
	fmt.Printf("%v \t %+v \t %#v\n", p1, p1, p1)
	fmt.Printf("%v \t %+v \t %#v\n", p2, p2, p2)

	fmt.Printf("\n==============================================\n\n")
	fmt.Println("指针类型接收者 可直接修改接收者中的值")
	p1.SetName("时间")
	fmt.Printf("%v \t %+v \t %#v\n", p1, p1, p1)

	fmt.Println("值类型接收者 不会直接修改接收者中的值 修改操作只针对副本有效")
	p1.SetName2("空间")
	fmt.Printf("%v \t %+v \t %#v\n", p1, p1, p1)
	p1 = p1.SetName2("空间")
	fmt.Printf("%v \t %+v \t %#v\n", p1, p1, p1)
}
