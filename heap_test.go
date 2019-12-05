package practice

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := NewMinHeap()

	h.Insert(3)
	h.Insert(6)
	h.Insert(1)
	h.Insert(5)
	h.Insert(2)
	h.Insert(4)

	if h.Size() != 6 {
		t.Errorf("MinHeap.Size() should be 6")
	}

	h.Insert(3)
	h.Insert(2)

	if v, ok := h.Poll(); !ok || v != 1 {
		t.Errorf("Lowest element should be 1")
	} else if h.Size() != 7 {
		t.Errorf("MinHeap.Size() should be 7")
	}

	if ok := h.Remove(6); !ok {
		t.Errorf("6 should have been found")
	} else if h.Size() != 6 {
		t.Errorf("MinHeap.Size() should be 6")
	}
}
