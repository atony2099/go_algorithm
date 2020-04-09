package sort

import (
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	type g struct{ name string }

	s := map[string]struct {
		input    int
		s        []int
		expected int
	}{
		"0": {0, []int{0, 1, 2}, 0},
		"1": {2, []int{0, 1, 2, 3, 4}, 2},
		"2": {3, []int{0, 1}, -1},
		"3": {3, []int{0}, -1},
	}

	for name, value := range s {
		t.Run(name, func(t *testing.T) {
			got := BinarySearch(value.s, value.input)

			if got != value.expected {
				t.Errorf("got:%d, expected:%d", got, value.expected)
			}
		})
	}

}

func TestBinarySearch2(t *testing.T) {

	type g struct{ name string }

	s := map[string]struct {
		input    int
		s        []int
		expected int
	}{
		"0": {0, []int{0, 1, 2}, 0},
		"1": {2, []int{0, 1, 2, 3, 4}, 2},
		"2": {3, []int{0, 1}, -1},
		"3": {3, []int{0}, -1},
	}

	for name, value := range s {
		t.Run(name, func(t *testing.T) {
			got := BinarySearch2(value.s, 0, len(value.s)-1, value.input)

			if got != value.expected {
				t.Errorf("got:%d, expected:%d", got, value.expected)
			}
		})
	}

}

func TestQuickSort(t *testing.T) {

	type g struct{ name string }

	s := map[string]struct {
		input    []int
		expected []int
	}{
		"0": {[]int{0, 3, 1}, []int{0, 1, 3}},
		"1": {[]int{0, 3, 1, 4}, []int{0, 1, 3, 4}},
		"2": {[]int{1, 0}, []int{0, 1}},
		"3": {[]int{0, 3, 1, 4, 4, 0, 2, 3}, []int{0, 0, 1, 2, 3, 3, 4, 4}},
	}

	for name, value := range s {
		t.Run(name, func(t *testing.T) {
			QuickSort(value.input, 0, len(value.input)-1)
			if reflect.DeepEqual(value.input, value.expected) == false {
				t.Errorf("got:%v, expected:%v", value.input, value.expected)
			}
		})
	}

}
