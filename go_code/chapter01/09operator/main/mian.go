package main

import "fmt"

func main() {

	fmt.Println(10 / 4)

	var n1 float32 = 10 / 4
	fmt.Println(n1)

	//希望保留小数部分，则需要有浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	//取余运算
	fmt.Println("10 % 3 = ", 10%3)     // 1
	fmt.Println("-10 % 3 = ", -10%3)   // -10 - (-10) / 3 * 3 = -1
	fmt.Println("-10 % -3 = ", 10%-3)  // 1
	fmt.Println("-10 % -3 = ", -10%-3) // -1

	//golang中, ++ 和 -- 只能独立使用！！！
	i := 1
	/*	错误示范:
		fmt.Println("", i++)
		var a int
		a = i++
		a = i--
		if i++ > 0 {
			fmt.Println("ok")
		}*/
	i++
	//++i 错误, golang中没有前++ 或 前--
	fmt.Println("", i)

	/***************逻辑运算符******************/
	//&& 逻辑与 也称短路与，如果第一个条件为false，则第二个条件不会判断，最终结果为false
	//|| 短路或同理 如果第一个条件为true，则第二个条件不会判断
	//全部都会输出
	if i > 0 && test() {
		fmt.Println("ok......")
	}
	//什么都不会输出，test()没有执行
	if i < 0 && test() {
		fmt.Println("ok......")
	}

	/***************赋值运算符******************/
	//有两个变量a b，要求将其进行交换不允许使用中间变量
	var a int = 10
	var b int = 20

	a = a + b
	b = a - b //等价于 b = a
	a = a - b //等价于 a = b

	fmt.Printf("a = %v b = %v", a, b)

	/***************优先级******************/
	//只有 单目运算符 赋值运算符 是从右向左运算的
	//其他都是从左向右
}

func test() bool {
	fmt.Println("test......")
	return true
}
