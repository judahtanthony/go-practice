package practice

import (
	"testing"
)

func TestAVL(t *testing.T) {
	collection := NewAVL()

	collection.Insert(1)
	collection.Insert(2)
	collection.Insert(3)
	collection.Insert(4)
	collection.Insert(5)

	if collection.Size() != 5 {
		t.Errorf("collection.Size() should be 5")
	}

	if !collection.Has(3) {
		t.Errorf("collection.Has(3) should be true")
	}

	if collection.Has(6) {
		t.Errorf("collection.Has(6) should be false")
	}

	collection.Remove(1)
	collection.Remove(3)
	collection.Remove(5)

	if collection.Size() != 2 {
		t.Errorf("collection.Size() should be 2")
	}

	if !collection.Has(2) {
		t.Errorf("collection.Has(2) should be true")
	}

	if collection.Has(1) {
		t.Errorf("collection.Has(1) should be false")
	}
}
