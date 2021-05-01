package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("j = ", j)
		}
	}
}
