package main

import "fmt"

/**
 * defer
 */
func sum(n1 int, n2 int) int {
	// 当执行到defer时，暂时不执行，会将defer后语句压入独立的栈中(defer栈(暂时这么称呼))
	// 当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1 =", n1)
	defer fmt.Println("ok1 n2 =", n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res =", res)
	return res
}

// 一般配合创建资源时，执行释放资源来用

func main() {
	res := sum(10, 20)
	fmt.Println("res =", res)
}
