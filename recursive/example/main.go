package main

import (
	"algorithm/recursive"
	"fmt"
)

func main() {

	// s := []int{1, 2, 3, 4}
	// r := recursive.GetSum(s)
	// fmt.Println(r)

	// res := recursive.Fibonacci(5)
	// fmt.Println(res)

	// res2 := recursive.DynamicFibonacciBottomUp(5)

	// fmt.Println(res2, "go")

	// res3 := recursive.DynamicFibonacciBottomUp(5)
	// fmt.Println(res3, "go")

	// res := recursive.MinCoins([]int{3}, 8)

	// fmt.Println(res)

	rr := recursive.MinCoinsBottomUp([]int{3, 2, 1}, 9)

	fmt.Println(rr)

}
