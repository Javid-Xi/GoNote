package main

import "fmt"

func main() {

	//string 基本使用
	var address string = "你好 Javid Xi!"
	fmt.Println(address)

	//字符串一旦赋值就不能修改了：在Go中字符串是不可变的
	var str = "hello"
	//str[0] = 'a' //这里不能修改str的内容，即go中字符串不可变
	fmt.Println(str)

	//字符串的两种表示形式
	//双引号:会识别转义字符
	//反引号: 以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击
	//输出源代码等效果
	str2 := "abc\nabc"
	fmt.Println(str2)

	str3 := `package main

import "fmt"

func main() {
	fmt.Println("hello world!")
}`

	fmt.Println(str3)

	//字符串拼接方式
	var str4 = "hello" + "world"
	str4 += " haha!"
	fmt.Println(str4)

	//当一个拼接操作很长时，怎么办，可以分行写
	//但注意要把符号保存到上一行
	var str5 = "hello " + "world " + "hello " + "world " +
		"hello " + "world " + "hello " + "world " +
		"hello " + "world "
	fmt.Println(str5)

	//基本数据类型的默认值
	var a int          // 0
	var b float32      // 0
	var c float64      // 0
	var isMarried bool //false
	var name string    // " "
	//%v 相应值的默认格式。
	fmt.Printf("a = %d, b = %f, c = %f, isMarried = %v, name = %v", a, b, c, isMarried, name)

}
