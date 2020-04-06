package main

import (
	"algorithm/lru"
	"fmt"
)

func main() {
	l := lru.Constructor(2)
	l.Put("1", 1)
	l.Put("2", 2)
	l.Put("3", 3)
	l.Put("4", 4)
	l.Link.Traverse()
	fmt.Println(l.Data)

}
