package recursive

import (
	"math"
)

// [1,2,3,4]
func GetSum(s []int) int {

	if len(s) == 1 {
		return s[0]
	}

	return s[0] + GetSum(s[1:])

}

func Fibonacci(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

// cache the result

func DynamicFibonacciBottomUp(n int) int {

	s := make([]int, n+1)
	s[0] = 0
	s[1] = 1

	for i := 2; i <= n; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return s[n]

}

// bottom up

var s = make([]int, 100)

func init() {

	for index := range s {
		s[index] = -1
	}

}

func DynamicFibonacciTopUp(n int) int {

	if s[n] == -1 {
		if n == 0 {
			s[0] = 0
		}
		if n == 1 {
			s[1] = 1
		}
		s[n] = DynamicFibonacciTopUp(n-1) + DynamicFibonacciTopUp(n-2)
	}

	return s[n]
}

func MinCoins(coins []int, change int) int {

	if change == 0 {
		return 0
	}
	res := math.MaxInt64
	for i := 0; i < len(coins); i++ {
		if change >= coins[i] {
			sub := MinCoins(coins, change-coins[i])
			if sub != math.MaxInt64 && sub+1 < res {
				res = sub + 1
			}
		}
	}
	return res

}

func MinCoinsBottomUp(coins []int, change int) int {
	changes := make(map[int]int, 100)

	for i := 0; i < 100; i++ {
		changes[i] = math.MaxInt64
	}
	changes[0] = 0
	for i := 1; i <= change; i++ {
		for _, coin := range coins {
			if coin <= i {
				sub, ok := changes[i-coin]
				if !ok {
					continue
				}
				if sub != math.MaxInt64 && sub+1 < changes[i] {
					changes[i] = sub + 1
				}

			}
		}
	}

	return changes[change]

}
