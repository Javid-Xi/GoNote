package exec

import "fmt"

func Exec01() {
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i) // 注意需要将i转成byte类型
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c\t", myChars[i])
	}
	fmt.Println()
}
