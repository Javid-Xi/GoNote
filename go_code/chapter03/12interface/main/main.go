package main

import "fmt"

/**
 * Description: 接口
 * 				golang 接口中的所有方法都没有方法体，即接口的方法都是没有实现的方法
				golang 中接口不需要显式地实现，只要一个变量，含有接口类型中的所有方法，
				那么这个变量就实现了这个接口，因此， golang没有implement这样的关键字

				注意事项:
				1. 一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋给该接口类型
				2. 接口中所有方法都没有方法体，即都是没有实现的方法
				3. 在golang中，一个自定义类型需要将某个接口的所有方法都实现，才说自定义类型实现了该方法
				4. 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
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
type integer int

type str string

func (i integer) Say() {
	fmt.Println("i = ", i)
}

func (s str) Say() {
	fmt.Println("s = ", s)
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

}
