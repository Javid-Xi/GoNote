package main

import "fmt"

func main() {
	//案例一
	var age int
	fmt.Println("请输入年龄: ")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你老了，兄弟~")
	} else { // 这里else不可换行
		fmt.Println("你还很年轻")
	}

	//案例二
	n1 := 10
	n2 := 20
	if n1+n2 > 25 {
		fmt.Println(n1 + n2)
	}

	//案例三
	n3 := 11.0
	n4 := 17.0
	if n3 > 10.0 && n4 < 20.0 {
		fmt.Println("和 =", n3+n4)
	}

	//if else if else
	var n int = 10
	if n > 9 {
		fmt.Println("ok1")
	} else if n > 6 {
		fmt.Println("ok2")
	} else if n > 3 {
		fmt.Println("ok3")
	} else {
		fmt.Println("ok4")
	}

}
