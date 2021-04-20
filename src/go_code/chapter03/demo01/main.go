package main

import "fmt"

//定义变量
func main() {
	//第一种：指定变量类型，声明后若不赋值，使用默认值
	//int 的默认值是0
	var i int
	i = 10
	fmt.Println("i =", i)

	//第二种：根据值自行判断变量类型（类型推导）
	var num = "tom"
	fmt.Println("num =", num)

	//第三种：省略var，注意 := 左侧的变量不应该是已经声明过的,否则会报错
	//下面方式等价于 var name string   name = "tom"
	name := "tom"
	fmt.Println("name =", name)
}
