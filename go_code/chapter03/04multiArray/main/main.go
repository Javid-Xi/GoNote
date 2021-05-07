package main

import "fmt"

/**
 * Description: 二维数组
 */
func main() {
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	fmt.Println(arr)

	// 遍历二维数组
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println("for-range 遍历方式: ")
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}

	fmt.Printf("arr address = %p\n", &arr)
	fmt.Printf("arr[0] address = %p\n", &arr[0])
	fmt.Printf("arr[0][0] address = %p\n", &arr[0][0])
	fmt.Printf("arr[1] address = %p\n", &arr[1])
	fmt.Printf("arr[1][0] address = %p\n", &arr[1][0])

}
