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
		if s[middle] > ele {
			right = middle - 1
		} else {
			left = middle + 1

		}
		middle = (left + right) / 2
	}
	return -1
}

func BinarySearch2(s []int, left, right, ele int) int {

	for left <= right {
		middle := (right + left) / 2

		if s[middle] == ele {
			return middle
		}
		if ele > s[middle] {
			return BinarySearch2(s, middle+1, right, ele)
		} else {
			return BinarySearch2(s, left, middle-1, ele)
		}
	}

	return -1

}

func partition(s []int, left, right int) int {
	pre := left
	pivot := s[left]

	for left < right {

		for s[right] >= pivot && left < right {
			right--
		}

		for s[left] <= pivot && left < right {
			left++
		}

		if left < right {
			s[left], s[right] = s[right], s[left]
		}

	}

	s[left], s[pre] = s[pre], s[left]

	return left

}

func QuickSort(s []int, left, right int) {
	if left < right {
		divide := partition(s, left, right)
		QuickSort(s, left, divide-1)
		QuickSort(s, divide+1, right)

	}
}
