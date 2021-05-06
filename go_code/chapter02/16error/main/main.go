package main

import (
	"errors"
	"fmt"
	"time"
)

/**
 * golang中的错误处理机制
 * 优点：加入预警代码提高代码健壮性
 */
func test() {
	// 使用defer + recovery 处理错误
	defer func() {
		// 内置函数可以捕获到异常
		if err := recover(); err != nil { // 说明捕获到错误
			fmt.Println("err =", err)
			// 这里可以将错误信息发送给管理员...
			fmt.Println("发送邮件给admin@163.com")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res =", res)
}

/**
 * Description: 自定义错误
 */

// 读取配置文件的信息
func readConfig(name string) (err error) {
	if name == "config.ini" {
		//读取
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误...")
	}
}

func test02() {
	if err := readConfig("config2.ini"); err != nil {
		panic(err)
	}
	fmt.Println("test02继续执行...")
}

func main() {

	// 测试错误处理机制
	test()
	for {
		fmt.Println("main下的方法")
		time.Sleep(time.Second)
	}

	// 测试自定义错误的使用
	//test02()
	//fmt.Println("main下的代码")
}
