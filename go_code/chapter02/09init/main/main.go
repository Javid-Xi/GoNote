package main

import (
	"fmt"
	"project01/go_code/chapter02/09init/utils"
)

// 每一个源文件都可以包含一个init函数
// 该函数会在main函数执行前被Go运行框架调用

var age = test()

func test() int {
	fmt.Println("test()...")
	return 90

}

// 通常在init函数中完成初始化工作

//执行顺序：全局变量定义 -> init函数 -> main函数
func init() {
	fmt.Println("init()...")
}

func main() {
	fmt.Println("main()...")
	fmt.Printf("Age = %v\n", utils.Age)
	fmt.Printf("Name = %v\n", utils.Name)
}
