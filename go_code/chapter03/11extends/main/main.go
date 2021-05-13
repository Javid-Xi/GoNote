package main

import "fmt"

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
	Student // 嵌入式 student 匿名结构体
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
}

type A struct {
	stu Student
	num int
}

type B struct {
	*Student
	num int
}
