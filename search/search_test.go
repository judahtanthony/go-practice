package search

import (
	"testing"
)

func TestLinearUnsorted(t *testing.T) {
	tests := []struct {
		Input1 []int
		Input2 int
		Expect bool
	}{
		{[]int{0, 1, 2, 3}, 2, true},
		{[]int{0, 1, 2}, 3, false},
		{[]int{}, 0, false},
	}
	for _, test := range tests {
		actual := LinearUnsorted(test.Input1, test.Input2)
		if test.Expect != actual {
			t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
		}
	}
}

func TestLinearSorted(t *testing.T) {
	tests := []struct {
		Input1 []int
		Input2 int
		Expect bool
	}{
		{[]int{0, 1, 2, 3}, 2, true},
		{[]int{0, 1, 2}, 3, false},
		{[]int{}, 0, false},
	}
	for _, test := range tests {
		actual := LinearSorted(test.Input1, test.Input2)
		if test.Expect != actual {
			t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
		}
	}
}

func TestBinary(t *testing.T) {
	tests := []struct {
		Input1 []int
		Input2 int
		Expect bool
	}{
		{[]int{0, 1, 2, 3}, 2, true},
		{[]int{0, 1, 2}, 3, false},
		{[]int{}, 0, false},
	}
	for _, test := range tests {
		actual := Binary(test.Input1, test.Input2)
		if test.Expect != actual {
			t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
		}
	}
}

func TestBinaryRecursive(t *testing.T) {
	tests := []struct {
		Input1 []int
		Input2 int
		Expect bool
	}{
		{[]int{0, 1, 2, 3}, 2, true},
		{[]int{0, 1, 2}, 3, false},
		{[]int{}, 0, false},
	}
	for _, test := range tests {
		actual := BinaryRecursive(test.Input1, test.Input2)
		if test.Expect != actual {
			t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
		}
	}
}
