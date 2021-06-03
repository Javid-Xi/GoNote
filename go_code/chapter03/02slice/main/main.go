package main

import (
	"GoNote/go_code/chapter03/02slice/exec"
	"fmt"
)

/**
 * Description: 切片 (可以简单理解成动态的数组)
 */

// 切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制
// 切片的使用和数组类似，遍历 访问 求长度等操作都一样
// 切片长度是可以变化的，因此切片可以是一个动态变化的数组
// 切片定义的基本语法:  var 切片名 []类型

func main() {
	// 切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}

	// 表示 slice 引用到intArr这个数组的 下标1 元素到 下标3 元素 (即 包头不包尾)
	slice := intArr[1:4]
	fmt.Println("intArr =", intArr)
	fmt.Println("slice 元素: ", slice)
	fmt.Println("slice len =", len(slice))
	fmt.Println("slice 的容量 =", cap(slice))

	// 切片内存布局分析
	fmt.Printf("intArr的地址 = %p\n", &intArr)
	fmt.Printf("intArr[1]的地址 = %p\n", &intArr[1])
	fmt.Printf("slice[0]的地址 = %p\n", &slice[0])
	fmt.Printf("slice的地址 = %p\n", &slice)

	slice[1] = 999
	fmt.Println(slice)
	fmt.Println(intArr)
	// 从上面可以看出:
	// slice的确是一个引用类型
	// 从底层来说，其实就是一struct结构体
	// type slice struct {
	// 		ptr *[2]int
	//      len int
	//		cap int
	// }

	/**
	 * Description: 切片的使用
	 */
	// 方式1 如上述...

	// 方式2 用过make来创建切片
	// 基本语法 var 切片名 []type = make([], len, [cap])
	var slicef []float64 = make([]float64, 5, 10)
	// 对于切片必须make后才能使用 否则为空
	slicef[1] = 10
	slicef[3] = 20
	fmt.Println(slicef)
	fmt.Println("slicef size =", len(slicef))
	fmt.Println("slicef cap =", cap(slicef))

	// 方式3 定义一个切片直接指定具体数组
	var slices []string = []string{"tom", "jack", "marry"}
	fmt.Println(slices)
	fmt.Println("slices size =", len(slices))
	fmt.Println("slices cap =", cap(slices))

	// 方式1和2的区别：
	// 方式1是直接引用数组，数组是事先存在的，程序员是可见的
	// 方式2是通过make来创建切片的，make也会创建一个数组，是由切片在底层进行维护，程序员不可见

	/**
	 * Description: 切片的遍历
	 */
	// 使用常规的for循环遍历切片
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d] = %d\t", i, slice[i])
	}
	fmt.Println()

	// 使用for-range 方式遍历切片
	for i, v := range slice {
		fmt.Printf("i = %v val = %v\t", i, v)
	}
	fmt.Println()

	/**
	 * Description: 切片注意事项和细节
	 */
	// slice := arr[0 : len(arr)] 可简写成 slice := arr[:]

	// 切片可以继续切片
	slice2 := slice[1:3]
	fmt.Println("slice2 =", slice2)

	slice2[1] = 100
	fmt.Println("slice2 =", slice2)
	fmt.Println("slice =", slice)
	fmt.Println("intArr =", intArr)

	// append 内置函数 可以对切片进行动态增加 用来扩容
	fmt.Println("--- append 内置函数 ---")
	var slice3 []int = intArr[:]
	fmt.Println("slice3 =", slice3)
	// 通过append给slice追加元素、
	//slice4 := append(slice3, 101)
	//fmt.Println("slice3 =", slice3)
	//fmt.Println("slice4 =", slice4)
	slice3 = append(slice3, 101)
	fmt.Println("slice3 =", slice3)

	slice3 = append(slice3, slice3...) // 要使用这种方式 后面必须是切片加三个点 不能是数组
	fmt.Println("slice3 =", slice3)

	// copy 切片的拷贝
	var slice4 = make([]int, 10)
	copy(slice4, slice3)
	fmt.Println("slice4 =", slice4)

	// slice3 和 slice4 数据空间是相互独立的
	slice4[2] = 0
	fmt.Println("slice3 =", slice3)
	fmt.Println("slice4 =", slice4)

	// 切片是引用类型 可以在函数修改对应内容
	test(slice4)
	fmt.Println("slice4 =", slice4)

	// 通过源代码发现，对于chan、map、slice等被当成指针处理，通过value.Pointer()获取对应的值的指针。
	// 当是slice类型的时候，返回是slice这个结构体里，字段Data第一个元素的地址。
	// 而对于len和cap 它们只是一个拷贝，如果要修改，那就要传递*slice作为参数才可以。

	// 练习
	fmt.Println(exec.Fibonacci(-10))
	fmt.Println(exec.Fibonacci(0))
	fmt.Println(exec.Fibonacci(1))
	fmt.Println(exec.Fibonacci(10))
}

func test(slice []int) {
	slice[0] = 888
}
