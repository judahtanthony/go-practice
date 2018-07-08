package practice

import (
	"testing"
)

func TestFibTD(t *testing.T) {
	for _, tt := range []struct {
		Input, Output int
	}{
		{0, 0},
		{1, 1},
		{10, 55},
		{20, 6765},
		//{50, 12586269025},
	} {
		if actual := FibTD(tt.Input); tt.Output != actual {
			t.Fatalf("expected: %v, actual: %v", tt.Output, actual)
		}
	}
}

func TestFibMemD(t *testing.T) {
	for _, tt := range []struct {
		Input, Output int
	}{
		{0, 0},
		{1, 1},
		{10, 55},
		{20, 6765},
		{50, 12586269025},
	} {
		if actual := FibMem(tt.Input); tt.Output != actual {
			t.Fatalf("expected: %v, actual: %v", tt.Output, actual)
		}
	}
}

func TestFibBU(t *testing.T) {
	for _, tt := range []struct {
		Input, Output int
	}{
		{0, 0},
		{1, 1},
		{10, 55},
		{20, 6765},
		{50, 12586269025},
	} {
		if actual := FibBU(tt.Input); tt.Output != actual {
			t.Fatalf("expected: %v, actual: %v", tt.Output, actual)
		}
	}
}
