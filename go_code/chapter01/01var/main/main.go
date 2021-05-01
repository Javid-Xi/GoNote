package main

import "fmt"

//定义变量
//定义全局变量
//var n1 = 100
//var n2 = 200
//var name = "jack"

//上面声明改为一次性声明
//var (
//	n3 = 300
//	n4 = 900
//	name = "mary"
//)

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

	/** 一次声明多个变量的 */
	//方式1
	//var n1, n2, n3 int
	//fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)

	//方式2
	//var n1, name, n3 = 100, "tom", 888
	//fmt.Println("n1=", n1, "name=", name, "n3=", n3)

	//方式3
	//使用类型推导
	n1, name, n3 := 100, "tom~", 888
	fmt.Println("n1=", n1, "name=", name, "n3=", n3)

	//输出全局变量
	//fmt.Println("n1=", n1, "name=", name, "n2=", n2)
	//fmt.Println("n3=", n3, "name=", name, "n4=", n4)
}
