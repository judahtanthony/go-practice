package practice

import (
	"reflect"
	"testing"
)

func TestCommonChars(t *testing.T) {
	tests := []struct {
		Input1 string
		Input2 string
		Expect string
	}{
		{"boo", "cat", ""},
		{"boo", "cot", "o"},
		{"book", "crabby", "b"},
		//{"book", "kangaroo", "ko"}, // This was my mistake
		{"book", "kangaroo", "koo"}, // This was my mistake
	}
	for _, test := range tests {
		res := CommonChars(test.Input1, test.Input2)
		actual := SortString(res)
		if test.Expect != actual {
			t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
		}
	}
}

func TestCommonChars2(t *testing.T) {
	tests := []struct {
		Input1 string
		Input2 string
		Expect string
	}{
		{"boo", "cat", ""},
		{"boo", "cot", "o"},
		{"book", "crabby", "b"},
		{"book", "kangaroo", "ko"},
	}
	for _, test := range tests {
		res := CommonChars2(test.Input1, test.Input2)
		actual := SortString(res)
		if test.Expect != actual {
			t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
		}
	}
}

func TestMapper(t *testing.T) {
	expected := make(map[int]int, 2)
	expected[2] = 1
	expected[4] = 2
	tests := []struct {
		Input1 []int
		Input2 func(int) int
		Expect map[int]int
	}{
		{[]int{1, 2}, func(x int) int { return x * 2 }, expected},
	}
	for _, test := range tests {
		actual := Mapper(test.Input1, test.Input2)
		if !reflect.DeepEqual(test.Expect, actual) {
			t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
		}
	}
}

func TestMapper2(t *testing.T) {
	expected := make(map[int]int, 2)
	expected[2] = 1
	expected[4] = 2
	tests := []struct {
		Input1 []int
		Input2 func(int) int
		Expect map[int]int
	}{
		{[]int{1, 2}, func(x int) int { return x * 2 }, expected},
	}
	for _, test := range tests {
		actual := Mapper2(test.Input1, test.Input2)
		if !reflect.DeepEqual(test.Expect, actual) {
			t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
		}
	}
}
