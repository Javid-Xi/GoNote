package exec

import "math"

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c *Circle) Area2() float64 {
	c.Radius = 20 // 传入指针会改变原数据的值
	return math.Pi * c.Radius * c.Radius
}
