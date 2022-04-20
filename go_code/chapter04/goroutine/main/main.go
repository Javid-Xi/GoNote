package main

import (
	"fmt"
	"sync"
	"time"
)

// go协程的特点：
// 有独立的栈空间
// 共享程序堆空间
// 调度由用户控制
// 协程是轻量级的线程

// M P G 模型
// M：操作系统的主线程（物理线程）
// P：协程执行需要的上下文
// G：协程

// 计算1-200各个数的阶乘
var (
	myMap = make(map[int]int, 10)
	// 声明一个全局互斥锁
	lock sync.Mutex
)

// 计算n的阶乘
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	// 开启多个协程执行任务（200个）
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 10)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]-%d\n", i, v)
	}
	lock.Unlock()
}
