package main

import "fmt"

func main()  {
	var i int32
	var j int32

	i, j = 1, 2

	i = i ^ j
	j = i ^ j
	i = i ^ j

	if i == 1 || j == 2 {
		fmt.Println("i and j are equal")
	}else {
		fmt.Println("i != j")
	}
}