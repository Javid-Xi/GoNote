package utils

import "fmt"

var Age int
var Name string

// Age 和 Name 全局变量，若需要在main.go使用
// 则需要初始化 Age 和 Name

// 被引入的包的 init  会先执行
func init() {
	Age = 18
	Name = "Javid"
	fmt.Println("utils init()...")
}
