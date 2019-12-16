package practice

import (
	"reflect"
	"testing"
)

func TestMergeArrays(t *testing.T) {
	tt := []struct {
		Input  [][]int
		Expect []int
	}{
		{
			[][]int{
				[]int{1, 3, 5},
				[]int{1, 3, 9},
				[]int{9, 5},
			},
			[]int{1, 3, 9, 5},
		},
	}
	for _, c := range tt {
		actual := MergeArrays(c.Input...)
		if !reflect.DeepEqual(actual, c.Expect) {
			t.Errorf("Expected: %v, Actual: %v", c.Expect, actual)
		}
	}
}
