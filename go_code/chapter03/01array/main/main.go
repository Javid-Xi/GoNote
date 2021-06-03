package main

import (
	"GoNote/go_code/chapter03/01array/exec"
	"fmt"
)

/**
 * Description: 数组的使用 遍历
 */
func main() {

	/**
	 * Description: 数组的使用
	 */
	// 定义完数组后，各个元素已经有默认值0了
	// 数组的地址可以通过数组名来获取 &intArr
	// 数组的第一个元素的地址，就是数组的首地址
	//var score [5]float64
	//for i := 0; i < len(score); i++ {
	//	fmt.Printf("请输入第%d个元素的\n", i + 1)
	//	fmt.Scanln(&score[i])
	//}
	//
	//// 变量数组打印
	//for i := 0; i < len(score); i++ {
	//	fmt.Printf("score[%d] = %v\n",i ,score[i])
	//}

	// 四种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println(numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Println(numArr02)

	var numArr03 = [...]int{5, 9, 5, 11, 2}
	fmt.Println(numArr03)

	var numArr04 = [...]int{1: 800, 0: 900, 3: 999}
	fmt.Println(numArr04)

	strArr := [...]string{1: "tom", 0: "jack", 2: "marry"}
	fmt.Println(strArr)

	/**
	 * Description: 数组的遍历
	 */
	// 常规遍历 略...

	// for-range结构遍历
	for i, v := range strArr {
		fmt.Printf("i = %v  v = %v\n", i, v)
	}
	for _, v := range strArr {
		fmt.Printf("元素的值 = %v\n", v)
	}

	// go的数组属于值类型， 在默认情况下是值拷贝， 因此进行会值拷贝，数组间不互相影响
	test01(strArr)
	fmt.Println(strArr)
	// 要改变原数组需要使用引用传递（指针方式）
	test02(&strArr)
	fmt.Println(strArr)

	/**
	 * Description: 数组练习
	 */
	// 练习一 输出A - Z
	exec.Exec01()

	// 练习二
	fmt.Printf("numArr03 数组中最大值为 %d\n", exec.Exec02(numArr03))

	// 练习三
	sum, aver := exec.Average(numArr03)
	fmt.Printf("数组numArr03 总和为 %d 平均值为 %f\n", sum, aver)

	// 练习四 生成随机长度为5的int数组 并将其反转
	// 生成随机数组
	nums := exec.RandArr()
	fmt.Println(nums)
	// 反转
	exec.Reverse(&nums)
	fmt.Println(nums)

}

func test01(arr [3]string) {
	arr[0] = "javid"
}
func test02(arr *[3]string) {
	(*arr)[0] = "javid"
}
