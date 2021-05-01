package main

import "fmt"

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'

	//当我们直接输出byte值，就是输出了对应的字符ASCII码值
	fmt.Println("c1 = ", c1)
	fmt.Println("c2 = ", c2)

	//如果要输出对应字符，需要使用格式化输出
	fmt.Printf("c1 = %c c2 = %c", c1, c2)

	//var c3 byte = '北' //overflow溢出
	//var c3 int = '北'
	c3 := '北'
	fmt.Printf("c3 = %c c3对应的码值 = %d", c3, c3)

}
