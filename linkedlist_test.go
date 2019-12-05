package practice

import (
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	sl := DoublyLinkedList{}

	if !sl.IsEmpty() {
		t.Errorf("List should be empty")
	}

	sl.Add(1)
	sl.Add(2)
	sl.Add(3)

	if sl.Size != 3 {
		t.Errorf("Size should be 3")
	}

	sl.AddFirst(0)
	sl.AddLast(4)

	if sl.Size != 5 {
		t.Errorf("Size should be 5")
	}

	if v, ok := sl.RemoveFirst(); ok == true && v != 0 {
		t.Errorf("First value should be 0")
	}

	if v, ok := sl.RemoveLast(); ok == true && v != 4 {
		t.Errorf("Last value should be 4")
	}

	for e := 1; e <= 3; e++ {
		if v, ok := sl.Remove(); ok == true && v != e {
			t.Errorf("Remove(): Expect %v, Actual: %v", e, v)
		}
	}
}
