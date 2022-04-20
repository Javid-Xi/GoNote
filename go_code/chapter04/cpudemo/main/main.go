package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum =", cpuNum)

	// 可以设置使用多少个cpu
	runtime.GOMAXPROCS(cpuNum - 2)
	fmt.Println("ok")
}
