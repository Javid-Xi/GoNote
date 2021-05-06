package main

import "fmt"

func main() {
	// len() 求长度

	// new 分配内存 给值类型分配
	num1 := 100
	fmt.Printf("num1的类型%T 值%v 地址%v\n", num1, num1, &num1)

	num2 := new(int)
	fmt.Printf("num2的类型%T num2的值%v num2的地址%v\n", num2, num2, &num2)
	fmt.Printf("默认指向的值是%v", *num2)

	// make 分配内存 分配引用类型内存

}
