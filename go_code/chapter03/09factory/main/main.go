package main

import (
	"fmt"
	"project01/go_code/chapter03/09factory/model"
)

func main() {
	//var stu1 = model.student{"Tom", 98.2}
	//fmt.Println(stu1)

	// 使用工厂模式 实现跨包
	var stu = model.NewStudent("Tom", 98.2)
	fmt.Println(*stu)
	fmt.Println(stu.GetName())
	fmt.Println(stu.GetScore())
}
