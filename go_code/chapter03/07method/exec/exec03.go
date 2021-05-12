package exec

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (stu *Student) String() string {
	return fmt.Sprintf("Name = %v\t Age = %v\n", stu.Name, stu.Age)
}
