package sort

func BinarySearch(s []int, ele int) int {

	//1. get middle index;
	left := 0
	right := len(s) - 1
	middle := len(s) / 2

	for left <= right {

		if s[middle] == ele {
			return middle
		}

		if left == right {
			if s[left] == ele {
				return left
			}
			return -1
		}

		if s[middle] > ele {
			right = middle - 1
		} else {
			left = middle + 1

		}
		middle = (left + right) / 2

	}

	return -1

}
