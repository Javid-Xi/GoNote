package main

import "fmt"

func main() {
	//goto可以无条件转移到程序指定的行
	//goto label
	//...
	//label: statement

	n := 20
	fmt.Println("ok1")
	if n >= 20 {
		goto label
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
label:
	fmt.Println("ok5")
	fmt.Println("ok6")

}
