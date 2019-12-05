package practice

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	h := MinHeap{make([]int, 0)}
	return &h
}

func (h *MinHeap) Size() int {
	return len(h.data)
}

func (h *MinHeap) IsEmpty() bool {
	return h.Size() == 0
}

func (h *MinHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.swim(h.Size() - 1)
}

func (h *MinHeap) Peek() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}
	return h.data[0], true
}

func (h *MinHeap) Poll() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}

	v := h.data[0]
	n := h.Size()

	h.swap(0, n-1)
	h.data = h.data[:n-1]
	h.sink(0)

	return v, true
}

func (h *MinHeap) Remove(v int) bool {
	if h.IsEmpty() {
		return false
	}

	for i, n := 0, h.Size(); i < n; i++ {
		if h.data[i] == v {
			if i < n-1 {
				h.swap(i, n-1)
				h.swim(h.sink(i))
			}
			h.data = h.data[:n-1]
			return true
		}
	}

	return false
}

func (h *MinHeap) swim(i int) int {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[i] >= h.data[parent] {
			break
		}
		h.swap(i, parent)
		i = parent
	}
	return i
}

func (h *MinHeap) sink(i int) int {
	n := h.Size() - 1
	for i <= n {
		left := i*2 + 1
		right := i*2 + 2
		if left > n {
			break
		}
		child := left
		if right <= n && h.data[right] < h.data[left] {
			child = right
		}
		if h.data[i] <= h.data[child] {
			break
		}
		h.swap(i, child)
		i = child
	}
	return i
}

func (h *MinHeap) swap(a, b int) {
	v := h.data[a]
	h.data[a] = h.data[b]
	h.data[b] = v
}
