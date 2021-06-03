package main

import (
	"GoNote/go_code/chapter03/07method/exec"
	"fmt"
)

/**
 * Description: 方法
 * 				func (receiver type) methodName (参数列表) (返回值列表) {
 *					方法体
 * 					return 返回值
 *				}
 */

type A struct {
	Num int
}

func (a A) test() {
	fmt.Println("test() a.Num =", a.Num)
	a.Num = 66
	fmt.Println("test() a.Num =", a.Num)
}

type Person struct {
	Name string
}

func (p Person) print() {
	fmt.Printf("%v是个好人\n", p.Name)
}

func (a A) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	// 结构体类型是值类型 在方法调用时 遵守值类型的传递机制 是值拷贝的传递方式
	// 如果希望在方法中修改结构体变量的值 可以通过结构体指针的方式来处理
	// Golang中的方法可以作用在指定的数据类型上
	var a A
	a.Num = 10
	a.test()
	fmt.Println("main() a.Num =", a.Num)

	var p Person
	p.Name = "Tom"
	p.print()

	fmt.Println(a.getSum(10, 20))

	// 练习一
	var c exec.Circle
	c.Radius = 10
	fmt.Println(c.Area())
	fmt.Println((&c).Area2())
	// 编译器做了优化当这里传地址时 使用 &c 等价于 c
	fmt.Println(c.Area2())
	fmt.Println(c.Area())

	//练习二
	var num exec.Integer = 10
	num.Print()
	num.Change()
	num.Print()

	// 练习三
	// 如果一个变量实现了String()这个方法 那么fmt.Println默认会调用这个变量的String()方法进行输出
	stu := exec.Student{
		Name: "Javid",
		Age:  22,
	}
	fmt.Println(stu)
	fmt.Println(&stu)
}
