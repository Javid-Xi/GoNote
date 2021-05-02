package main

import (
	"fmt"
	"strings"
)

/**
 * 闭包 : 一个函数与其相关的引用环境组合的一个整体(实体)
 */
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += "a" + string(36)
		fmt.Println(str)
		return n
	}
}

// 小练习
// 返回的匿名函数引用到的的外部变量构成一个闭包
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		// 如果 name 没有指定后缀则加上，否则返回原来的名字
		//HasPrefix 判断字符串 s 是否以 prefix 开头：
		//HasSuffix 判断字符串 s 是否以 suffix 结尾
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 使用普通函数实现上述同样功能
func makeSuffix2(suffix string, name string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}

func main() {
	//  累加器
	f := AddUpper()
	fmt.Println(f(1)) // 11
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16

	// 闭包的说明: 9 - 13 行就是一个闭包
	// 返回的是一个匿名函数，但是这个匿名函数引用到函数外的 n,
	// 因此这个匿名函数就和 n 形成一个整体，构成闭包。

	// 小练习
	// 返回一个闭包
	f2 := makeSuffix(".jpg")
	fmt.Println("文件处理后 : ", f2("winter"))
	fmt.Println("文件处理后 : ", f2("brid.jpg"))

	// 使用普通函数实现同样功能
	fmt.Println("文件处理后 : ", makeSuffix2(".jpg", "brid"))

	// 结论: 使用闭包只需要传一次后缀变量即可 仔细体会
}
