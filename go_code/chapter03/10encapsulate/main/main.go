package main

import (
	"GoNote/go_code/chapter03/10encapsulate/model"
	"fmt"
)

func main() {

	account1 := model.NewAccount("javid1012", "999999", 40)
	if account1 != nil {
		fmt.Println("账户创建成功: ", *account1)
	} else {
		fmt.Println("账户创建失败")
	}

	account1.SetAccount("5641325")
	fmt.Println(*account1)

	account1.SetAccount("564125")
	fmt.Println(*account1)

}
