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

func test01(n1 float64, n2 float64) {
	var tmp float64 = n1
	n1 = n2
	n2 = tmp
}

func test02(n1 *float64, n2 *float64) {
	var tmp float64 = *n1
	*n1 = *n2
	*n2 = tmp
}

//go函数不支持函数重载
//go语言中可变参数和空接口完全可以实现此效果
/*func test02(n1 *int)  {
	var tmp int = *n1
	*n1 = *n2
	*n2 = tmp
}*/

func getSum(n1 float64, n2 float64) float64 {
	return n1 + n2
}

//函数作形参
func myFun(funvar func(float64, float64) float64, num1 float64, num2 float64) float64 {
	return funvar(num1, num2)
}

//自定义数据类型
type myInt int
type myFumType func(float64, float64) float64

func myFun2(funvar myFumType, num1 float64, num2 float64) float64 {
	return funvar(num1, num2)

}

//支持对函数返回值命名
func getSumAndSub(n1 float64, n2 float64) (sum float64, sub float64) {
	sub = n1 - n2
	sum = n1 + n2
	return
}

//go中支持可变参数
//args是slice切片 通过args[index] 可以访问到各个值
//如果一个函数的形参列表中有可变参数，则可变参数需放在形参列表最后
//零到多个参数
func sum1(args ...int) {

}

//1到多个参数
func sum2(n1 int, args ...int) int {
	sum := n1
	//遍历args
	for i := 0; i < len(args); i++ {
		sum += args[i] //args[i] 表示取出args切片的第i个值
	}
	return sum
}

func main() {
	var x float64
	var y float64
	fmt.Scanf("%f %f", &x, &y)

	fmt.Printf("x + y = %f\n", cal(x, y, '+'))
	fmt.Printf("x - y = %f\n", cal(x, y, '-'))
	fmt.Printf("x * y = %f\n", cal(x, y, '*'))
	fmt.Printf("x / y = %f\n", cal(x, y, '/'))

	//在golang中 基本数据类型和数组都是值传递
	fmt.Println("交换 x y : ")
	fmt.Printf("x = %v  y = %v\n", x, y)
	test01(x, y)
	fmt.Printf("x = %v  y = %v\n", x, y)
	test02(&x, &y)
	fmt.Printf("x = %v  y = %v\n", x, y)

	//函数也是一种数据类型
	//可以赋值给一个变量 则该变量就是一个函数类型的变量了
	//通过改变量可以实现对函数的调用
	a := getSum
	fmt.Printf("a的数据类型%T, getSum的类型是%T\n", a, getSum)
	res := a(x, y)
	fmt.Println("res = ", res)

	res2 := myFun(getSum, x, y)
	fmt.Println("res2 = ", res2)

	//自定义数据类型
	var num1 myInt
	num1 = 40
	//var num2 int = num1 错误
	fmt.Println("num1 = ", num1)

	res3 := myFun2(getSum, x, y)
	fmt.Println("res3 = ", res3)

	//支持对函数返回值命名
	sum, sub := getSumAndSub(x, y)
	fmt.Printf("sum = %v  sub = %v\n", sum, sub)

	//使用_标识符，忽略返回值
	sum, _ = getSumAndSub(x, y)

	//go中支持可变参数
	fmt.Println("测试可变参数的使用: ")
	fmt.Println("传一个参数: ", sum2(1))
	fmt.Println("传多个参数: ", sum2(1, 2, 3, 4, 5, 6))

}
