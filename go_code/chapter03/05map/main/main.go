package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/**
 * Description: map
 *				类似其它语言中的哈希表或字典，以key-value形式存储数据
 *				key必须是支持==或!=比较运算的类型，不可以是函数、map或slice
 *				Map通过key查找value比线性搜索快很多
 *				Map使用make()创建，支持:=这种简写方式
 *
 *				make([keyType]valueType,cap),cap表示容量，可省略
 *				超出容量时会自动扩容，但尽量提供一个合理的初始值
 *				使用len()获取元素个数
 *				键值对不存在时自动添加，使用delete()删除某键值对
 *				使用for range对map和slice进行迭代
 */

func main() {
	// key 通常是 int 或 string
	// 注意： slice map function 不可以，因为这几个无法使用 == 来判断

	var a map[string]string
	// 声明不会分配内存 需要使用make来分配内存
	a = make(map[string]string, 3)
	a["no1"] = "Javid"
	a["no2"] = "Marry"
	a["no3"] = "Tom"
	a["no4"] = "Tom"
	a["no5"] = "Tom"
	fmt.Println(a)
	// 随便测试下 map 的扩容
	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano()) // 给其当前的时间戳 给纳秒比较好
		a[strconv.FormatInt(int64(rand.Intn(1000)), 10)] = "Javid" + strconv.FormatInt(int64(rand.Intn(1000)), 10)
		time.Sleep(time.Millisecond) // 这里若不使用延时将导致添加不进去元素 因为 map 的扩容需要时间
	}
	fmt.Println("a size =", len(a))
	fmt.Println(a)

	a["no2"] = "Jack"
	fmt.Println(a)

	// map的使用
	// 方式一 如前所述
	// 方式二
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"
	fmt.Println(cities)

	// 第三种方式
	var heroes map[string]string = map[string]string{
		"no1": "宋江",
		"no2": "卢俊义",
		"no3": "吴用",
	}
	heroes["no4"] = "林冲"
	fmt.Println(heroes)

	// 课堂案例
	studentsMap := make(map[string]map[string]string)

	studentsMap["stu01"] = make(map[string]string, 3)
	studentsMap["stu01"]["name"] = "tom"
	studentsMap["stu01"]["sex"] = "男"
	studentsMap["stu01"]["address"] = "北京"

	studentsMap["stu02"] = make(map[string]string, 3)
	studentsMap["stu02"]["name"] = "marry"
	studentsMap["stu02"]["sex"] = "女"
	studentsMap["stu02"]["address"] = "上海"

	fmt.Println(studentsMap)

	/**
	 * Description: map 的增删改查
	 */
	delete(cities, "no1")
	fmt.Println(cities)

	// 当delete指定的key不存在时，删除就不会操作
	delete(cities, "no5")
	fmt.Println(cities)

	// map的清空
	// golang 无清空map的内置函数
	// 1. 遍历删除
	// 2. 直接make一个新空间
	cities = make(map[string]string)
	fmt.Println(cities)

	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"

	// map的查找
	val, ok := cities["no1"]
	if ok {
		fmt.Printf("有no1 key 值为 %v\n", val)
	} else {
		fmt.Println("没有 no1 的key")
	}

	/**
	 * Description: map 的遍历
	 */
	// 只能用for-range
	for key, v := range cities {
		fmt.Printf("key = %v\tvalue = %v\n", key, v)
	}

	// 结构较复杂的 map 遍历
	for k1, v1 := range studentsMap {
		fmt.Println("k1 =", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\tk2 = %v\tv2 = %v\n", k2, v2)
		}
	}

	// map 的长度
	fmt.Println("studentsMap size =", len(studentsMap))
	fmt.Println("cities size =", len(cities))

	/**
	 * Description: map 的切片
	 */
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"
	}
	fmt.Println(monsters)

	// 要加入元素
	// 需要使用切片的append函数 可以动态增加 monster切片
	newMonster := map[string]string{
		"name": "火云邪神",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
}
