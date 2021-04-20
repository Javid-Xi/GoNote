package main

import "fmt" //提供格式化、输入、输出的函数

func main() {
	fmt.Println("tom\tyack")
	fmt.Println("hello\nworld")
	fmt.Println("hello\nworld")
	fmt.Println("E:\\Study\\golang\\Go_Learning\\src\\go_code\\chapter01\\main") //第一个斜杠表示转义
	fmt.Println("tom说\"i love you")
	fmt.Println("天龙八部雪山飞狐狸\r张飞") //\r 回车 从当前行最前开始输出，覆盖掉以前内容
	//fmt.Println("天龙八部\n张飞")
}
