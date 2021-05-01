package main

import "fmt"

//演示golang中的基本数据类型的转换
//必须都要用强制转换
func main() {
	var i int = 100
	//希望将 i => float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)

	fmt.Printf("i = %v n1 = %v n2 = %v n3 = %v\n", i, n1, n2, n3)

	//转换溢出
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Printf("num2 = %v", num2)

}
