package practice

import (
	"reflect"
	"testing"
)

func TestFindDuplicatesBrute(t *testing.T) {
	tests := []struct {
		Input  []int
		Expect []int
	}{
		{[]int{0, 1, 1, 3}, []int{1}},
		{[]int{0, 0, 2, 2}, []int{0, 2}},
		{[]int{0, 1, 1, 1, 2, 2}, []int{1, 1, 1, 2}},
		{[]int{0, 1, 2, 3}, []int{}},
		{[]int{}, []int{}},
	}
	for _, test := range tests {
		actual := FindDuplicatesBrute(test.Input)
		if !reflect.DeepEqual(test.Expect, actual) {
			t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
		}
	}
}

func TestFindDuplicatesSort(t *testing.T) {
	tests := []struct {
		Input  []int
		Expect []int
	}{
		{[]int{0, 1, 1, 3}, []int{1}},
		{[]int{0, 0, 2, 2}, []int{0, 2}},
		{[]int{0, 1, 1, 1, 2, 2}, []int{1, 1, 2}},
		{[]int{0, 1, 2, 3}, []int{}},
		{[]int{}, []int{}},
	}
	for _, test := range tests {
		actual := FindDuplicatesSort(test.Input)
		if !reflect.DeepEqual(test.Expect, actual) {
			t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
		}
	}
}

func TestFindDuplicatesSortFind(t *testing.T) {
	tests := []struct {
		Input  []int
		Expect []int
	}{
		{[]int{0, 1, 1, 3}, []int{1}},
		{[]int{0, 0, 2, 2}, []int{0, 2}},
		{[]int{0, 1, 1, 1, 2, 2}, []int{1, 1, 2}},
		{[]int{0, 1, 2, 3}, []int{}},
		{[]int{}, []int{}},
	}
	for _, test := range tests {
		actual := FindDuplicatesSortFind(test.Input)
		if !reflect.DeepEqual(test.Expect, actual) {
			t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
		}
	}
}

func TestFindDuplicatesCount(t *testing.T) {
	tests := []struct {
		Input  []int
		Expect []int
	}{
		{[]int{0, 1, 1, 3}, []int{1}},
		{[]int{0, 0, 2, 2}, []int{0, 2}},
		{[]int{0, 1, 1, 1, 2, 2}, []int{1, 1, 2}},
		{[]int{0, 1, 2, 3}, []int{}},
		{[]int{}, []int{}},
	}
	for _, test := range tests {
		actual := FindDuplicatesCount(test.Input, 3)
		if !reflect.DeepEqual(test.Expect, actual) {
			t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
		}
	}
}

func TestFindMaxBrute(t *testing.T) {
	tests := []struct {
		Input       []int
		Expect      int
		ExpectError bool
	}{
		{[]int{0, 1, 2, 3}, 0, false},
		{[]int{0, 1, 1, 2}, 1, false},
		{[]int{0, 1, 1, 2, 2}, 1, false},
		{[]int{7}, 7, false},
		{[]int{}, 0, true},
	}
	for _, test := range tests {
		actual, err := FindMaxBrute(test.Input)
		if test.ExpectError {
			if err == nil {
				t.Errorf("Expected: %v to produce error\n", test.Input)
			}
		} else {
			if err != nil {
				t.Errorf("%v shoud NOT produce error\n", test.Input)
			} else if test.Expect != actual {
				t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
			}
		}
	}
}

func TestFindMaxSort(t *testing.T) {
	tests := []struct {
		Input       []int
		Expect      int
		ExpectError bool
	}{
		{[]int{0, 1, 2, 3}, 0, false},
		{[]int{0, 1, 1, 2}, 1, false},
		{[]int{0, 1, 1, 2, 2}, 1, false},
		{[]int{7}, 7, false},
		{[]int{}, 0, true},
	}
	for _, test := range tests {
		actual, err := FindMaxSort(test.Input)
		if test.ExpectError {
			if err == nil {
				t.Errorf("Expected: %v to produce error\n", test.Input)
			}
		} else {
			if err != nil {
				t.Errorf("%v shoud NOT produce error\n", test.Input)
			} else if test.Expect != actual {
				t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
			}
		}
	}
}

func TestFindMaxCount(t *testing.T) {
	tests := []struct {
		Input       []int
		Expect      int
		ExpectError bool
	}{
		{[]int{0, 1, 2, 3}, 0, false},
		{[]int{0, 1, 1, 2}, 1, false},
		{[]int{0, 1, 1, 2, 2}, 1, false},
		{[]int{7}, 7, false},
		{[]int{}, 0, true},
	}
	for _, test := range tests {
		actual, err := FindMaxCount(test.Input, 3)
		if test.ExpectError {
			if err == nil {
				t.Errorf("Expected: %v to produce error\n", test.Input)
			}
		} else {
			if err != nil {
				t.Errorf("%v shoud NOT produce error\n", test.Input)
			} else if test.Expect != actual {
				t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
			}
		}
	}
}

func TestFindMajorityCancelation(t *testing.T) {
	tests := []struct {
		Input    []int
		Expect   int
		ExpectOk bool
	}{
		{[]int{7}, 7, true},
		{[]int{0, 1, 1}, 1, true},
		{[]int{0, 1, 1, 2, 2, 2, 2}, 2, true},
		{[]int{0, 1}, 0, false},
		{[]int{0, 1, 1, 2, 2, 2}, 0, false},
		{[]int{}, 0, false},
	}
	for _, test := range tests {
		actual, ok := FindMajorityCancelation(test.Input)
		if !test.ExpectOk {
			if ok {
				t.Errorf("%v should NOT return majority\n", test.Input)
			}
		} else {
			if !ok {
				t.Errorf("%v shoud return majority\n", test.Input)
			} else if test.Expect != actual {
				t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
			}
		}
	}
}

func TestFindMissingCount(t *testing.T) {
	tests := []struct {
		Input    []int
		Expect   int
		ExpectOk bool
	}{
		{[]int{1}, 2, true},
		{[]int{2}, 1, true},
		{[]int{2, 3}, 1, true},
		{[]int{1, 3}, 2, true},
		{[]int{1, 2}, 3, true},
		{[]int{1, 2, 3, 4, 5, 7}, 6, true},
		{[]int{}, 0, false},
	}
	for _, test := range tests {
		actual, ok := FindMissingCount(test.Input)
		if !test.ExpectOk {
			if ok {
				t.Errorf("%v should NOT return majority\n", test.Input)
			}
		} else {
			if !ok {
				t.Errorf("%v shoud return majority\n", test.Input)
			} else if test.Expect != actual {
				t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
			}
		}
	}
}

func TestFindMissingSum(t *testing.T) {
	tests := []struct {
		Input    []int
		Expect   int
		ExpectOk bool
	}{
		{[]int{1}, 2, true},
		{[]int{2}, 1, true},
		{[]int{2, 3}, 1, true},
		{[]int{1, 3}, 2, true},
		{[]int{1, 2}, 3, true},
		{[]int{1, 2, 3, 4, 5, 7}, 6, true},
		{[]int{}, 0, false},
	}
	for _, test := range tests {
		actual, ok := FindMissingSum(test.Input)
		if !test.ExpectOk {
			if ok {
				t.Errorf("%v should NOT return majority\n", test.Input)
			}
		} else {
			if !ok {
				t.Errorf("%v shoud return majority\n", test.Input)
			} else if test.Expect != actual {
				t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
			}
		}
	}
}

func TestFindMissingXOR(t *testing.T) {
	tests := []struct {
		Input    []int
		Expect   int
		ExpectOk bool
	}{
		{[]int{1}, 2, true},
		{[]int{2}, 1, true},
		{[]int{2, 3}, 1, true},
		{[]int{1, 3}, 2, true},
		{[]int{1, 2}, 3, true},
		{[]int{1, 2, 3, 4, 5, 7}, 6, true},
		{[]int{}, 0, false},
	}
	for _, test := range tests {
		actual, ok := FindMissingXOR(test.Input)
		if !test.ExpectOk {
			if ok {
				t.Errorf("%v should NOT return majority\n", test.Input)
			}
		} else {
			if !ok {
				t.Errorf("%v shoud return majority\n", test.Input)
			} else if test.Expect != actual {
				t.Errorf("Expeced: %v, Actual: %v\n", test.Expect, actual)
			}
		}
	}
}

func TestFindSumTwoSort(t *testing.T) {
	tests := []struct {
		Input1   []int
		Input2   int
		Expect1  int
		Expect2  int
		ExpectOk bool
	}{
		{[]int{1, 1}, 2, 1, 1, true},
		{[]int{1, 2, 3, 4}, 4, 1, 3, true},
		{[]int{4, 3, 2, 1}, 4, 1, 3, true},
		{[]int{}, 0, 0, 0, false},
		{[]int{1}, 1, 0, 0, false},
	}
	for _, test := range tests {
		actual1, actual2, ok := FindSumTwoSort(test.Input1, test.Input2)
		if (test.ExpectOk && !ok) || (!test.ExpectOk && ok) {
			not := ""
			if !test.ExpectOk {
				not = " NOT"
			}
			t.Errorf("%v with %v should%s return OK\n", test.Input1, test.Input2, not)
		} else if (test.Expect1 != actual1) || (test.Expect2 != actual2) {
			t.Errorf(
				"Expected: %v & %v, Actual: %v & %v\n",
				test.Expect1, test.Expect2, actual1, actual2)
		}
	}
}

func TestFindSumTwoMap(t *testing.T) {
	tests := []struct {
		Input1   []int
		Input2   int
		Expect1  int
		Expect2  int
		ExpectOk bool
	}{
		{[]int{1, 1}, 2, 1, 1, true},
		{[]int{1, 2, 3, 4}, 4, 1, 3, true},
		{[]int{4, 3, 2, 1}, 4, 1, 3, true},
		{[]int{}, 0, 0, 0, false},
		{[]int{1}, 1, 0, 0, false},
	}
	for _, test := range tests {
		actual1, actual2, ok := FindSumTwoMap(test.Input1, test.Input2)
		if (test.ExpectOk && !ok) || (!test.ExpectOk && ok) {
			not := ""
			if !test.ExpectOk {
				not = " NOT"
			}
			t.Errorf("%v with %v should%s return OK\n", test.Input1, test.Input2, not)
		} else if (test.Expect1 != actual1) || (test.Expect2 != actual2) {
			t.Errorf(
				"Expected: %v & %v, Actual: %v & %v\n",
				test.Expect1, test.Expect2, actual1, actual2)
		}
	}
}

func TestFindMinAbsSum(t *testing.T) {
	tests := []struct {
		Input    []int
		Expect1  int
		Expect2  int
		ExpectOk bool
	}{
		{[]int{1, 1}, 1, 1, true},
		{[]int{1, 0, -12}, 0, 1, true},
		{[]int{-5, 0, 1}, 0, 1, true},
		{[]int{-5, -1, 0, 3, 10}, -1, 0, true},
		{[]int{}, 0, 0, false},
		{[]int{1}, 0, 0, false},
	}
	for _, test := range tests {
		actual1, actual2, ok := FindMinAbsSum(test.Input)
		if (test.ExpectOk && !ok) || (!test.ExpectOk && ok) {
			not := ""
			if !test.ExpectOk {
				not = " NOT"
			}
			t.Errorf("%v should%s return OK\n", test.Input, not)
		} else if (test.Expect1 != actual1) || (test.Expect2 != actual2) {
			t.Errorf(
				"Expected: %v & %v, Actual: %v & %v\n",
				test.Expect1, test.Expect2, actual1, actual2)
		}
	}
}

func TestFindBitonicMaxBSearch(t *testing.T) {
	tests := []struct {
		Input    []int
		Expect   int
		ExpectOk bool
	}{
		{[]int{0, 1, 0}, 1, true},
		{[]int{0, 1, 1, 0}, 1, true},
		{[]int{0, 1, 5, 3, 0}, 5, true},
		{[]int{0, 1, 4, 6, 0}, 6, true},
		{[]int{0, 1, 4, 6, 6, 0}, 6, true},
		{[]int{}, 0, false},
		{[]int{0}, 0, false},
		{[]int{0, 1}, 0, false},
		{[]int{0, 1, 2}, 0, false},
	}
	for _, test := range tests {
		actual, ok := FindBitonicMaxBSearch(test.Input)
		if (test.ExpectOk && !ok) || (!test.ExpectOk && ok) {
			not := ""
			if !test.ExpectOk {
				not = " NOT"
			}
			t.Errorf("%v should%s be OK.", test.Input, not)
		} else if test.Expect != actual {
			t.Errorf("Input: %v\nExpected: %v\nActual: %v", test.Input, test.Expect, actual)
		}
	}
}
