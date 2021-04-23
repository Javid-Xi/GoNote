package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string //空的str

	fmt.Println("第一种方式：")
	//使用第一种方式来转换 fmt.Sprintf
	//Sprintf根据format参数生成格式化的字符串并返回该字符串。
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str = %q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str = %q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str = %q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str = %q\n", str, str)

	fmt.Println("第二种方式：")
	//第二种方式 strconv方式
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str = %q\n", str, str)

	//说明：'f'是格式 10表示小数位保留10位 64表示这个小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str = %q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str = %q\n", str, str)
}
