package practice

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf, _ := NewUnionFind(10)
	// |0, 1, 2, 3, 4, 5, 6, 7 ,8, 9|
	// |0, 1, 2, 3, 4, 5, 6, 7 ,8, 9|
	uf.Union(0, 5)
	// |0, 1, 2, 3, 4, 0, 6, 7 ,8, 9|
	// |0, 1, 2, 3, 4, 5, 6, 7 ,8, 9|
	uf.Union(1, 7)
	// |0, 1, 2, 3, 4, 0, 6, 1 ,8, 9|
	// |0, 1, 2, 3, 4, 5, 6, 7 ,8, 9|
	uf.Union(2, 7)
	// |0, 2, 2, 3, 4, 0, 6, 1 ,8, 9|
	// |0, 1, 2, 3, 4, 5, 6, 7 ,8, 9|
	uf.Union(4, 5)
	// |4, 2, 2, 3, 4, 4, 6, 1 ,8, 9|
	// |0, 1, 2, 3, 4, 5, 6, 7 ,8, 9|
	uf.Union(7, 9)
	// |4, 2, 2, 3, 4, 4, 6, 2 ,8, 2|
	// |0, 1, 2, 3, 4, 5, 6, 7 ,8, 9|

	if conn, _ := uf.Connected(0, 1); conn {
		t.Errorf("0 should NOT be connected to 1")
	}
	if conn, _ := uf.Connected(1, 2); !conn {
		t.Errorf("1 should be connected to 2")
	}
}
