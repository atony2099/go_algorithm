package sort

import (
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
		// "1": {2, []int{0, 1, 2, 3, 4}, 2},
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
