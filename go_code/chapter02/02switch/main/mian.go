package main

import (
	"fmt"
)

func test(b byte) byte {
	return b + 1
}

func main() {
	var key byte
	fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)
	switch test(key) + 1 {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
		//...
	default:
		fmt.Println("输入有误")
	}

	//switch 后也可以不带表达式，类似if -- else 分支来使用
	var score int = 66
	switch {
	case score > 90:
		fmt.Println("成绩优秀")
	case score >= 70 && score <= 90:
		fmt.Println("成绩优良")
	case score >= 60 && score < 70:
		fmt.Println("成绩及格")
	default:
		fmt.Println("不及格")
	}

	//switch 后也可以直接声明/定义一个变量，分号结束，不推荐
	switch grade := 90; {
	case grade > 90:
		fmt.Println("成绩优秀")
	case grade >= 70 && grade <= 90:
		fmt.Println("成绩优良")
	case grade >= 60 && grade < 70:
		fmt.Println("成绩及格")
	default:
		fmt.Println("不及格")
	}

	//switch 穿透 fallthrought
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到...")
	}
	//不建议使用穿透，
	//如果希望 10 和 20 执行同样的逻辑 可以这样写
	switch num {
	case 10, 20:
		fmt.Println("ok1")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到...")
	}

	//switch 语句还可以被用于使用于 type - switch 类判断某个 interface 变量中实际指向的变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) { //可以判断x这个空接口究竟指向那种数据类型
	case nil:
		fmt.Printf("x 的类型~ : %T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
