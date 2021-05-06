package main

import "fmt"

func main() {

	// string 底层是一个byte数组 因此string也可以进行切片处理
	str := "hello Javid"
	// 使用切片获得 Javid
	slice := str[6:]
	fmt.Println(slice)

	// string 是不变的不可通过 str[0] = 'z' 修改其值
	// 先将string -> []byte 再 []rune -> 重写转成string

	arr1 := []byte(str) // 先转成byte切片
	fmt.Println(&str)
	fmt.Println(&arr1[0])
	arr1[0] = 'z'
	str = string(arr1) // 修改完后再转为字符串
	fmt.Println(str)

	// 细节: 转成 []byte 可以处理英文和数字，但不能处理中文
	// 将 string 转为 []rune 即可，因为 []rune 是按字符处理的
	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println(str)
}
