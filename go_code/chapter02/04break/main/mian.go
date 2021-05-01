package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break //默认跳出最近的for循环
			}
			fmt.Println("j = ", j)
		}
	}

	fmt.Println("测试标签: ")
lable2: //设置一个标签 名字随意
	for i := 0; i < 4; i++ {
		//lable1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				break lable2 //跳出标签对应的for循环
			}
			fmt.Println("j = ", j)
		}
	}
}
