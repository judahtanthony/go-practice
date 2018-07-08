package practice

import (
	"reflect"
	"sort"
	"testing"
)

func TestPerm(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Output []string
	}{
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{
			"abc",
			"acb",
			"bac",
			"bca",
			"cab",
			"cba",
		}},
	} {
		actual := Perm(tt.Input)
		sort.Strings(actual)
		if !reflect.DeepEqual(tt.Output, actual) {
			t.Fatalf("Expected: %v, Actual: %v", tt.Output, actual)
		}
	}
}

func TestPermChan(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Output []string
	}{
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{
			"abc",
			"acb",
			"bac",
			"bca",
			"cab",
			"cba",
		}},
	} {
		actual := []string{}
		for str := range PermChan(tt.Input) {
			actual = append(actual, str)
		}
		sort.Strings(actual)
		if !reflect.DeepEqual(tt.Output, actual) {
			t.Fatalf("Expected: %v, Actual: %v", tt.Output, actual)
		}
	}
}
