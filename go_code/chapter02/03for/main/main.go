package main

import (
	"fmt"
)

func main() {
	//第一种写法
	for i := 1; i <= 3; i++ {
		fmt.Println("hello Javid")
	}

	//for循环的第二种写法
	j := 1
	for j <= 3 {
		fmt.Println("hello Javid Xi")
		j++
	}

	//for循环的第三种写法
	//死循环, 通常需要配合break
	k := 1
	for {
		fmt.Println("Javid loop")
		if k >= 3 {
			break
		}
		k++
	}

	/***************** 字符串遍历****************/
	//方式一
	//字符串含有中文会出现乱码
	//因为 传统字符串是按照字节来遍历 而一个汉字在utf8编码对应3个字节
	var str string = "Javid Xi 西安"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	fmt.Println()

	//只需要将 str 转成 []rune 切片即可
	str2 := []rune(str) //把 str 转成 []rune
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c ", str2[i])
	}

	//方式二 推荐使用这种方式
	//这种方式是按照字符来遍历的
	str = "javid.cn 上海"
	for index, val := range str {
		fmt.Printf("\nindex = %d, val = %c ", index, val)
	}
	fmt.Println()

	level := 10
	//打印金字塔
	for i := 0; i < level; i++ {
		for j := 0; j < level-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= 2*i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//打印空心金字塔
	for i := 0; i < level; i++ {
		for j := 0; j < level-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= 2*i; j++ {
			if i == level-1 || j == 0 || j == 2*i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	//打印乘法表
	// i 表层数
	var num int = 20
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}
