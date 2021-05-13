package main

import "fmt"

type Student struct {
	Name   string
	Gender string
	Age    int
	Id     string
	Score  float64
}

func (stu *Student) say() string {
	infoStr := fmt.Sprintf("name = %v\tgender = %v\tage = %v\tid = %v\t"+
		"score = %v", stu.Name, stu.Gender, stu.Age, stu.Id, stu.Score)
	return infoStr
}

func main() {
	var stu = Student{
		Name:   "tom",
		Gender: "男",
		Age:    18,
		Id:     "01",
		Score:  97.9,
	}
	fmt.Println(stu.say())

	stu2 := Student{"mary", "女", 18, "02", 98.2}
	fmt.Println(stu2.say())

	// stu3是个指针
	stu3 := &Student{"Jack", "女", 18, "02", 98.2}
	fmt.Println(stu3)
}
