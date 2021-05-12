package main

import (
	"encoding/json"
	"fmt"
)

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

	/**
	 * Description: 结构体内存分配机制
	 */
	fmt.Println("--- 结构体内存分配机制 ---")
	var p4 *Person = &p1
	p4.Name = "Javid"
	p4.Age = 22
	fmt.Printf("p1.Name = %v\tp1.Name = %v\n", p1.Name, p1.Name)
	fmt.Printf("p4.Name = %v\tp4.Name = %v\n", p4.Name, p4.Name)

	fmt.Printf("p1的地址 = %p\n", &p1)
	fmt.Printf("p4的地址 = %p\n", &p4)
	fmt.Printf("p4的值 = %p\n", p4)

	/**
	 * Description: 结构体使用细节
	 */
	// 结构体所有字段在内存中连续分布
	// 结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段

	// 结构体进行type重新定义（相当于取别名） golang认为是新的数据类型 但是相互可以强转
	var p5 Per
	// p5 = p1 错误 需要强制转换
	p5 = Per(p1)
	fmt.Println(p5)

	// struct 的每个字段上，可以写一个tag，该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化
	monster := Monster{"牛魔王", 500, "芭蕉扇~"}

	// 将monster变量序列化成json字符串
	// json.Marshal 中使用到反射 （反射章节会有详细介绍 ）
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("jsonStr = ", string(jsonStr))
}

type Per Person

type Person struct {
	Name string
	Age  int
}

type Monster struct {
	Name  string `json:"name"` // 结构体标签 struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}
