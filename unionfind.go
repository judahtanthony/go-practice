package practice

import (
	"fmt"
)

type UnionFind struct {
	data []int
}

func NewUnionFind(size int) (*UnionFind, error) {
	if size < 0 {
		return nil, fmt.Errorf("size must be a postive integer")
	}
	uf := UnionFind{make([]int, size)}
	for i := 0; i < size; i++ {
		uf.data[i] = i
	}
	return &uf, nil
}

func (uf *UnionFind) Connected(a, b int) (bool, bool) {
	rootA, okA := uf.Find(a)
	rootB, okB := uf.Find(b)

	return rootA == rootB, okA && okB
}
func (uf *UnionFind) Union(a, b int) bool {
	rootA, okA := uf.Find(a)
	rootB, okB := uf.Find(b)

	if !okA || !okB {
		return false
	}

	// I'm skipping the component size check for brevity.
	if rootA != rootB {
		uf.data[rootB] = rootA
	}

	return true
}

func (uf *UnionFind) Find(n int) (int, bool) {
	if n >= len(uf.data) {
		return 0, false
	}

	root := n
	for root != uf.data[root] {
		root = uf.data[root]
	}

	for n != root {
		uf.data[n], n = root, uf.data[n]
	}

	return root, true
}

func (uf *UnionFind) FindR(n int) (int, bool) {
	if n >= len(uf.data) {
		return 0, false
	}
	if n == uf.data[n] {
		return n, true
	}

	root, ok := uf.FindR(uf.data[n])
	if ok {
		uf.data[n] = root
	}

	return root, ok
}
