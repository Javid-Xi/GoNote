package exec

import "fmt"

type Integer int

func (i Integer) Print() {
	fmt.Println("i =", i)
}

func (i *Integer) Change() {
	*i += 1
}
