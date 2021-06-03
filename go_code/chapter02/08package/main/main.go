package main

import (
	"GoNote/go_code/chapter02/08package/exec"
	"GoNote/go_code/chapter02/08package/utils"
	util "GoNote/go_code/chapter02/08package/utils2" //给包取别名
	"fmt"
)

//编译时需要编译main包所在的文件夹
//项目的目录结构最好按照规范来组织
//编译后生成一个有默认名的可执行文件，在$GOPATH目录下，可以指定名字和目录，
//例如：编译的指令: go build -o bin/my.exe project01/go_code/chapter02/08package/main

func main() {
	var x float64
	var y float64
	fmt.Println("请输入x y: ")
	fmt.Scanf("%f %f", &x, &y)

	fmt.Printf("x + y = %.2f\n", utils.Cal(x, y, '+'))
	fmt.Printf("x - y = %.2f\n", utils.Cal(x, y, '-'))
	fmt.Printf("x * y = %.2f\n", utils.Cal(x, y, '*'))
	fmt.Printf("x / y = %.2f\n", utils.Cal(x, y, '/'))
	fmt.Printf("~~ x / y = %.2f\n", util.Cal(x, y, '/'))
	for i := 1; i <= 10; i++ {
		fmt.Print(exec.FibonacciNum(i), " ")
	}
}
