package main

import "fmt"

func main() {
	// 演示管道
	// 1. 创建一个可以存3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)
	// 2. 看看intChan是什么类型
	fmt.Printf("intChan的值=%v inChan本身的地址=%p\n", intChan, &intChan)

	// 3. 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	// intChan <- num
	// intChan <- num

	// 注意⚠️：给管道写入的数据不能超过其容量
	// 否则会报错：fatal error: all goroutines are asleep - deadlock!

	// 4. 看看管道的长度和Cap（容量）
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))

	// 5. 从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2 =", num2)

	// 在没有使用协程情况下，若管道数据已完全取出，再取则会报错（错误同上）
	num3 := <-intChan
	// num4 := <- intChan
	fmt.Printf("num3 = %d\n", num3)
	//fmt.Printf("num4 = %d\n", num4)
}
