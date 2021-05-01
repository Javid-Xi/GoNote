package main

import "fmt"

// func 函数名 (形参列表) (返回值列表) {
//	执行语句
//  return 返回值列表
//}
func cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号错误...")
	}
	return res
}

func main() {
	var x float64
	var y float64
	fmt.Scanf("%f %f", &x, &y)

	fmt.Printf("x + y = %f\n", cal(x, y, '+'))
	fmt.Printf("x - y = %f\n", cal(x, y, '-'))
	fmt.Printf("x * y = %f\n", cal(x, y, '*'))
	fmt.Printf("x / y = %f\n", cal(x, y, '/'))
}
