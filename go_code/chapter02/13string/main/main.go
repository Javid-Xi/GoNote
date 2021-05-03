package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * 字符串中常用的系统函数 (重点)
 */

func main() {

	// 统计字符串长度 len(str) //内建函数 builtin
	str := "Javid"
	fmt.Println("str len =", len(str))
	str2 := "北京" // golang的编码统一为utf-8，一个汉字占用三个字节
	fmt.Println("str len =", len(str2))

	// 字符串遍历，同时处理中文的问题 r := []rune(str)
	for i := 0; i < len(str); i++ {
		fmt.Printf("字符 = %c\t", str2[i])
	}
	fmt.Println()
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符 = %c\t", r[i])
	}
	fmt.Println()

	// 字符串转整数 n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("12")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果是", n)
	}

	// 整数转字符串 str = strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("str = %v\n", str)

	// 字符串转成byte切片
	var bytes = []byte("hello go")
	fmt.Printf("byte = %c\n", bytes) // 这里%v对应其ascii编码

	// []byte 转 string
	str = string(bytes)
	fmt.Println(str)

	// 十进制 转 2 8 16进制 strconv.FormatInt(i, base) 返回对应的字符串
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的 16 进制是: %v\n", str)
	fmt.Printf("123对应的 2 进制是: %v\n", strconv.FormatInt(123, 2))

	// 查找子串是否在指定的字符串中
	b := strings.Contains("Javid hello", "vid")
	fmt.Println(b)

	// 统计一个字符串中有几个指定的子串
	num := strings.Count("cehseesehese", "se")
	fmt.Println(num)

	// 不区分大小写的字符串比较 ( == 是区分字母大小写的)
	b = strings.EqualFold("abc", "AbC")
	fmt.Println(b, "abc" == "AbC")

	// 返回子串在字符串中第一次出现的位置，若没有则返回 -1
	index := strings.Index("NLT_abc", "abc")
	fmt.Println("index =", index)
	index = strings.Index("NLT_abc", "hel")
	fmt.Println("index =", index)

	// 返回子串在字符串最后一次出现的位置
	index = strings.LastIndex("go golang", "go")
	fmt.Println("index =", index)

	// 将指定子串替换成另外一个子串
	str2 = "go go go go hello"
	str = strings.Replace(str2, "go", "go语言", 1)
	fmt.Println(str)
	str = strings.Replace(str2, "go", "go语言", 2)
	fmt.Println(str)
	// -1为全换
	str = strings.Replace(str2, "go", "go语言", -1)
	fmt.Println(str)
	fmt.Println("str2 本身无变化 :", str2)

	// 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
	strArr := strings.Split(str2, " ")
	fmt.Printf("%v\n", strArr)
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v] = %v\n", i, strArr[i])
	}

	// 将字符串的字母进行大小写的转换
	str = "goLANG HELLO"
	str = strings.ToLower(str)
	fmt.Println(str)

	// 将字符串左右两边空格去掉
	str = "   Golang Hello "
	fmt.Println(str)
	fmt.Println(strings.TrimSpace(str))

	// 将字符串左右两边指定的字符串去掉
	str = " !! ! ??! hello go ? ?! !! "
	fmt.Println(str)
	fmt.Println(strings.Trim(str, " !?"))

	// 将字符串左边指定字符去掉
	// 将字符串右边指定字符去掉
	fmt.Println(strings.TrimLeft(str, " !?"))
	fmt.Println(strings.TrimRight(str, " !"))

	str = "Javid Xi"
	// 判断字符串是否以指定的字符串开头
	fmt.Println(strings.HasPrefix(str, "Jav"))
	// 判断字符串是否以指定的字符串结尾
	fmt.Println(strings.HasSuffix(str, "Xi"))
}
