package model

import "fmt"

type account struct {
	accountNo string
	pwd       string
	balance   float64
}

// NewAccount 工厂模式的函数-构造函数
func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长度输入错误...")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码长度输入错误...")
		return nil
	}
	if balance < 20 {
		fmt.Println("余额不足...")
		return nil
	}
	return &account{accountNo, pwd, balance}
}

func (a *account) SetAccount(accountNo string) {
	if len(accountNo) != 6 {
		fmt.Println("密码格式错误")
	} else {
		a.accountNo = accountNo
	}
}
