package main

import (
	"fmt"
	"math/rand"
	"project01/go_code/chapter03/12interface/exec"
	"sort"
	"time"
)

/**
 * Description: 接口
 				接口是引用类型
 				golang 接口中的所有方法都没有方法体，即接口的方法都是没有实现的方法
 				golang 中接口不需要显式地实现，只要一个变量，含有接口类型中的所有方法，
				那么这个变量就实现了这个接口，因此， golang没有implement这样的关键字

				注意事项:
				1. 接口本身不能实现创建实例，但是可以指向一个实现了该接口的自定义类型的变量(实例)
				2. 一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋给该接口类型
				3. 接口中所有方法都没有方法体，即都是没有实现的方法
				4. 在golang中，一个自定义类型需要将某个接口的所有方法都实现，才说自定义类型实现了该方法
				5. 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
				6. 一个自定义类型可以实现多个接口
				7. golang接口中不能有任何变量
				8. 一个接口可以继承多个别个接口， 此时要实现A接口也必须实现B、C接口

				接口是基于方法的
				接口继承的接口中不能有相同的方法
*/

// 声明一个接口

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作......")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作......")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作......")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作......")
}

type Computer struct {
}

// 编写一个working方法 接受一个Usb接口类型变量
// 只要是实现了Usb接口(所谓Usb接口就是指实现了Usb接口声明的所有方法)

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}
type integer int

type str string

func (i integer) Say() {
	fmt.Println("i = ", i)
}

func (s str) Say() {
	fmt.Println("s = ", s)
}
func (s str) Hello() {
	fmt.Println("hello!", s)
}

func main() {

	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)
	// 松耦合高内聚

	// 注意事项 4
	var i integer = 10
	var s str = "Javid"
	var a AInterface = i
	a.Say()
	a = s
	a.Say()

	var b BInterface = s
	b.Hello()

	// 空接口可以接受任何一种数据类型
	var t1 T = i
	fmt.Println(t1)
	var t2 interface{} = s
	fmt.Println(t2)

	// 接口经典案例 实现Interface接口 实现可掉用库函数排序方法
	var heroes exec.HeroSlice
	rand.Seed(time.Now().UnixNano()) // 注入当前的时间戳 随机数按照传入的数来生成 不传入则生成的值永远不变
	for i := 0; i < 10; i++ {
		hero := exec.Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(60),
		}
		heroes = append(heroes, hero)
	}

	fmt.Println("排序前:")
	fmt.Println(heroes)
	sort.Sort(heroes)
	fmt.Println("按年龄排序后:")
	fmt.Println(heroes)

}

type T interface {
}
