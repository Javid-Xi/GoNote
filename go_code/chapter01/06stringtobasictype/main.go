package main

import (
	"fmt"
	"strconv"
)

//string 转基本数据类型
func main() {
	var str string = "true"
	var b bool
	//此函数有两个返回值 (value bool, err error)
	//用下划线忽略掉第二个返回值
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b typr %T b = %v\n", b, b)

	//转失败，默认为false
	b2, _ := strconv.ParseBool("45")
	fmt.Printf("b2 typr %T b2 = %v\n", b2, b2)

	var str2 string = "1234567890"
	var n1 int64
	//base为进制
	//ParseInt返回值为int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 := int(n1)
	fmt.Printf("n1 type %T n1 = %v\n", n1, n1)
	fmt.Printf("n2 typr %T n2 = %v\n", n2, n2)

	var str3 string = "12.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("n2 typr %T n2 = %v\n", f1, f1)

	//若字符串不是全数字，转换成int默认为0
	f2, _ := strconv.ParseFloat("11hello", 64)
	fmt.Printf("n2 typr %T n2 = %v\n", f2, f2)
}
