package main

import (
	"algorithm/linkedList/single"
	"fmt"
)

func main() {

	l := map[int]int{1: 1, 2: 2, 3: 3}
	// delete

	fmt.Println(len(l), "---")

	list := single.Single{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// list.Traverse()
	v := list.GetValueAtIndex(1)
	fmt.Println(v)

	// fmt.Println(list.DeletValue(2))
	// fmt.Println(list.DeletValue(3))
	// fmt.Println(list.DeletValue(1))

	list.Reverse()
	// fmt.Println(list.Header, list.Tail)
	list.Traverse()

	// fmt.Println(list.Tail, list.Length)

}
