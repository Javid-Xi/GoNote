package main

import (
	"fmt"
)

/**
 * Description: 多重继承
 * 如何理解go语言提倡组合，不提倡继承 https://www.jianshu.com/p/150523db21a9
 */

type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) showInfo() {
	fmt.Printf("学生名%v 年龄%v 成绩%v\n", stu.Name, stu.Age, stu.Score)

}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

type Pupil struct {
	Student // 嵌入student 匿名结构体
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中...")
}

type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试中...")
}

type E struct {
	Student
	int
}

func main() {

	pupil := &Pupil{}
	// pupil.Student.Name = "Tom"
	// pupil.Student.Age = 9
	// 简化
	pupil.Name = "Tom"
	pupil.Age = 9
	pupil.testing()
	pupil.SetScore(95)
	pupil.showInfo()

	graduate := &Graduate{}
	graduate.Name = "mary"
	graduate.Age = 18
	graduate.testing()
	graduate.SetScore(79)
	graduate.showInfo()

	// 结构体可以使用嵌套匿名结构体中的所有字段和方法

	// 就近访问原则

	// 结构体嵌入两个或多个匿名结构体，若两个结构体有相同的字段和方法时，
	// 且结构体本身没有同名的字段和方法，在访问时，就必须指定匿名结构体名字

	// 若结构体嵌套了一个有名结构体，这种模式就是组合，在访问组合的结构体的字段和方法时，必须带上该结构体名称
	var a = A{Student{"Tom", 18, 98}, 1}
	// a.Age = 18 错误
	a.stu.Age = 18
	fmt.Println(a)

	var b = B{&Student{"Jack", 19, 95}, 2}
	fmt.Println(*b.Student, b.num)

	// 匿名字段是基本数据类型
	var e E
	e.Name = "狐狸精"
	e.Age = 300
	e.int = 20
	fmt.Println(e)
	// 如果一个struct嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现多重继承
	// 注：尽量不要使用多重继承

	// 组合相对于继承的优点在于
	// 1. 可以利用面向接口编程原则的一系列优点，封装性好，耦合性低
	// 2. 相对于继承的编译期确定实现，组合的运行态指定实现，更加灵活
	// 3. 组合是非侵入式的，继承是侵入式的

	// 嵌套匿名结构体后，也可以在创建结构体变量时，直接指定各个匿名结构体字段的值
	tv := TV{
		Goods{
			Price: 5000.99,
			Name:  "电视机002",
		},
		Brand{
			Name:    "夏普",
			Address: "北京",
		},
	}
	// 访问 tv 成员
	// fmt.Println(tv.Name) 错误
	fmt.Println(tv.Goods.Name)
	fmt.Println(tv.Price)
	fmt.Println(tv.Brand.Name)
	fmt.Println(tv.Address)
}

type A struct {
	stu Student
	num int
}

type B struct {
	*Student
	num int
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}
