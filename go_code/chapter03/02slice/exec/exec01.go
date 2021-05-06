package exec

import (
	"fmt"
)

func Fibonacci(n int) []uint64 {
	defer func() {
		if err := recover(); err != nil { // 说明捕获到错误
			fmt.Println("err =", err)
		}
	}()

	// 声明一个切片 切片大小为 n
	fbnSlice := make([]uint64, n)

	if n == 0 {
		return fbnSlice
	}

	fbnSlice[0] = 1
	if n == 1 {
		return fbnSlice
	}
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}
