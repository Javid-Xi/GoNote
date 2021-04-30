package main

import "fmt"

//演示golang中的指针类型
func main() {
	var i int = 10
	// i 的地址
	fmt.Println("i的地址: ", &i)

	//ptr是一个指针变量
	//ptr的类型为*int
	//ptr本身值为&i
	var ptr *int = &i
	fmt.Printf("ptr = %v\n", ptr)
	fmt.Println("ptr的地址: ", &ptr)
	fmt.Println("ptr指向地址的值为: ", *ptr)

	num := 9
	fmt.Printf("num = %v\n", num)
	ptr = &num
	*ptr = 10
	fmt.Printf("num = %v\n", num)

}

/**
 * 值类型：基本数据类型 int系列 float系列 bool string 数组 结构体struct 通常在栈区
 * 引用类型：指针、size切片、map、管道chan、interface 等都是引用类型 通常在堆区
 */
