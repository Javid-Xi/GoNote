package main

import "fmt"

/**
 * 匿名函数
 */

// 全局匿名函数
var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 用法一
	// 在定义匿名函数时就直接调用 这种方式匿名函数只能使用一次
	// 求两个数的和
	n1 := 10
	n2 := 20
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(n1, n2)
	fmt.Println("res =", res)

	//用法二
	//将匿名函数赋给变量
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(n1, n2)
	fmt.Println("res2 =", res2)

	//用法三
	//全局匿名函数
	res3 := Fun1(n1, n2)
	fmt.Println("res3 =", res3)
}
