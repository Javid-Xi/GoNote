package main

import (
	"fmt"
	"math/rand"
	"sort"
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

	/**
	 * Description: map 排序
	 */
	// go 中没有一个专门的方法针对map的 key 进行排序
	// go 中 map默认是无序的
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[12] = 100
	map1[1] = 120
	map1[2] = 20
	map1[4] = 40
	map1[8] = 80
	fmt.Println(map1)

	// 利用切片排序
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	// 排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Printf("map1[%v] = %v\n", k, map1[k])
	}

	/**
	 * Description: map 的使用细节和陷阱
	 */
	// map是引用类型 作为参数传到其他函数里的修改会影响原来map的值
	fmt.Println("--- map 的使用细节和陷阱 ---")
	modify(map1)
	fmt.Println("map1 =", map1)

	// map会自动进行扩容 而切片不会
	// map的 value 经常使用结构体类型， 更适合管理复杂的数据结构
	// map的key为学生学号 value为学生结构的结构体
	studentsMap2 := make(map[string]Stu, 10)
	// 创建两个学生
	stu1 := Stu{"tom", 18, "北京"}
	stu2 := Stu{"mary", 20, "上海"}
	studentsMap2["no1"] = stu1
	studentsMap2["no2"] = stu2
	fmt.Println("studentsMap2 =", studentsMap2)

	// 遍历方式打印
	for k, v := range studentsMap2 {
		fmt.Printf("学号%v\t", k)
		fmt.Printf("姓名%v\t", v.Name)
		fmt.Printf("年龄%v\t", v.Age)
		fmt.Printf("地址%v\n", v.Address)
	}

	/**
	 * Description: 练习
	 */
	// 如果某个用户名存在就将其密码修改为"888888"
	// 若不存在就增加这个用户信息
	fmt.Println("--- 练习 ---")
	users := make(map[string]map[string]string, 10)
	users["smith"] = make(map[string]string, 2)
	users["smith"]["pwd"] = "777777"
	users["smith"]["nickname"] = "小花猫"
	fmt.Println("原users =", users)
	modifyUser(users, "tom")
	modifyUser(users, "mary")
	modifyUser(users, "smith")
	fmt.Println("users =", users)

}

// 定义学生结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}

func modify(map1 map[int]int) {
	map1[10] = 999
}

func modifyUser(users map[string]map[string]string, name string) {
	// 判断users中有无name
	if users[name] != nil {
		// 有这个用户
		users[name]["pwd"] = "888888"
	} else {
		// 没有这个用户
		users[name] = make(map[string]string)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称~" + name // 示意
	}
}
