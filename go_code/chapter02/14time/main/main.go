package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
 * 时间和日期相关函数
 */
func main() {
	// time.Time类型
	// 获取当前时间
	now := time.Now()
	fmt.Printf("now = %v\nnow type = %T\n", now, now)

	// 通过now可以获取到年月日，时分秒
	fmt.Printf("年 = %v\n", now.Year())
	fmt.Printf("月 = %v\n", now.Month())
	fmt.Printf("月 = %v月\n", int(now.Month()))
	fmt.Printf("日 = %v\n", now.Day())
	fmt.Printf("时 = %v\n", now.Hour())
	fmt.Printf("分 = %v\n", now.Minute())
	fmt.Printf("秒 = %v\n", now.Second())

	// 格式化 日期时间
	// 方式1 使用Sprintf 或 Printf
	fmt.Printf("当前时间是的是: %d-%d-%d %d:%d:%d\n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("当前时间是的是: %d-%d-%d %d:%d:%d\n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr = %v\n", dateStr)

	// 方式二 使用time.Format()方法
	fmt.Printf(now.Format("2006/01/02 15:04:05")) // 格式只能用此时间来表达
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15时04分05秒"))
	fmt.Println()

	// 时间常量
	// Nanosecond 纳秒 Microsecond 微秒 Millisecond 毫秒
	// Second 秒 Minute 分钟 Hour 小时
	i := 0
	// 每隔1秒打印
	for {
		i++
		fmt.Println(i)
		// 休眠
		time.Sleep(time.Second)

		if i == 3 {
			break
		}
	}

	// 每隔0.1秒打印
	i = 1
	for ; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}

	// Unix 和 UnixNano
	// 表示unix时间 即从时间点 January 1 1970 UTC 到 时间点t 所经过的时间(Unix 单位为秒 UnixNano为毫秒)
	// 其作用可以用于获取随机数字
	fmt.Printf("unix时间戳 =%v\nunixnano时间戳 = %v\n", now.Unix(), now.UnixNano())

	// 小练习 获取函数执行时间
	start := time.Now().Unix()
	test()
	end := time.Now().Unix()
	fmt.Printf("执行test()所用时间为: %v秒\n", end-start)

}

func test() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
