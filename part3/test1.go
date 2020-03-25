package main 

import "fmt"

type Integer int 

func (a Integer) Less(b Integer) bool {
	return a < b
}

func main()  {
	var b Integer = 1
	if b.Less(2) {
		fmt.Println(b, "Less 2")
	}
}