package dynamic

func ClimbStairs(n int) int {
	if n < 1 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return ClimbStairs(n-1) + ClimbStairs(n-2)
}

// func ClimbStairs2(n int) int {
// 	if n  {

// 	}

// }
