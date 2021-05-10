package main

import "fmt"

/**
 * Description: struct
 *				go中的struct与其他编程语言的class有同等地位
 * 				go 面向对象是基于 struct 的
 *				go 中面向接口编程 是非常重要的特性
 * 				go 中面向对象编程非常简洁，去掉传统OOP语言的继承、方法重载、构造函数、析构函数、隐藏this指针等等
 */

// 定义一个Cat结构体，将Cat的各个字段信息，放入Cat结构体中进行管理
type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	// 快速入门 面向对象方式解决养猫问题
	// 创建一个Cat变量
	var cat1 Cat
	fmt.Println("结构体变量cat1默认值: ", cat1)
	cat1.Name = "小白"
	cat1.Age = 4
	cat1.Color = "red"
	fmt.Println("cat1: ", cat1)

	// go 中的结构体是值类型的 是直接指向实例空间的
	// 结构体字段 = 属性= field

	// 创建结构体实例的四种方式
	// 方式一 直接声明

	// 方式二 最常使用
	var p1 = Person{"mary", 20}
	fmt.Println(p1)

	// 方式三
	var p2 *Person = new(Person)
	// p2 是一个指针
	(*p2).Name = "smith"
	(*p2).Age = 19

	fmt.Println(*p2)
	// 赋值也可这样写
	// go设计值为了程序员使用方便 底层会对下面语句进行处理自动加上取值运算 *
	p2.Age = 24
	fmt.Println(*p2)

	// 方案四
	var p3 *Person = &Person{"Javid", 22}
	fmt.Println(*p3)

}

type Person struct {
	Name string
	Age  int
}
