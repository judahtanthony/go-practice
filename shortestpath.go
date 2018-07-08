package practice

import (
	"container/heap"
	"errors"
	"math"
)

type Node struct {
	Cost  int
	Edges []Edge
}

type Edge struct {
	Node int
	Cost int
}

// An Item is something we manage in a priority queue.
type Item struct {
	Node     int
	NodeCost int
	Cost     int
	Via      int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Cost+pq[i].NodeCost < pq[j].Cost+pq[j].NodeCost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, cost int, via int) {
	item.Cost = cost
	item.Via = via
	heap.Fix(pq, item.index)
}

func Dijkstra(nodes [][]Edge, start int, end int) ([]int, error) {
	n := len(nodes)
	if n < 1 {
		return nil, errors.New("Must pass non-empty slice.")
	}
	if start < 0 || start >= n {
		return nil, errors.New("Invalid arguments: start is out of range")
	}
	if end < 0 || end >= n {
		return nil, errors.New("Invalid arguments: end is out of range")
	}
	if start == end {
		return []int{start}, nil
	}

	// I'm sure there is a better way to access previous
	// items, but I just don't know it.  I with this heap
	// had an efficient "contains" or "get" like a map.
	visited := make([]*Item, n)
	queue := make(PriorityQueue, n)

	// Fill the queue with all our nodes.
	for i := range nodes {
		item := &Item{
			Node:  i,
			Cost:  math.MaxInt64,
			Via:   i,
			index: i,
		}
		visited[i] = item
		queue[i] = item
	}
	// Start here.
	visited[start].Cost = 0
	heap.Init(&queue)

	var curr *Item
	for {
		curr = heap.Pop(&queue).(*Item)
		// We found it!
		if curr.Node == end {
			break
		}
		// Not it.  Let's check this nodes edges.
		for _, edge := range nodes[curr.Node] {
			item := visited[edge.Node]
			cost := curr.Cost + edge.Cost
			// Did we find a less expensive way to get here?
			if cost < item.Cost {
				queue.update(item, cost, curr.Node)
			}
		}
		// We failed to find the path.
		if queue.Len() == 0 {
			return nil, errors.New("Failed to find path")
		}
	}
	// Gather the path.
	path := []int{}
	for curr != nil {
		path = append(path, curr.Node)
		if curr.Node == curr.Via {
			break
		}
		curr = visited[curr.Via]
	}
	// We have to reverse the path.
	n = len(path)
	for i := 0; i < n/2; i++ {
		path[i], path[n-i-1] = path[n-i-1], path[i]
	}
	return path, nil
}

func AStar(nodes []Node, start int, end int) ([]int, error) {
	n := len(nodes)
	if n < 1 {
		return nil, errors.New("Must pass non-empty slice.")
	}
	if start < 0 || start >= n {
		return nil, errors.New("Invalid arguments: start is out of range")
	}
	if end < 0 || end >= n {
		return nil, errors.New("Invalid arguments: end is out of range")
	}
	if start == end {
		return []int{start}, nil
	}

	// I'm sure there is a better way to access previous
	// items, but I just don't know it.  I with this heap
	// had an efficient "contains" or "get" like a map.
	visited := make([]*Item, n)
	queue := make(PriorityQueue, n)

	// Fill the queue with all our nodes.
	for i, node := range nodes {
		item := &Item{
			Node:     i,
			NodeCost: node.Cost,
			Cost:     math.MaxInt64 - node.Cost,
			Via:      i,
			index:    i,
		}
		visited[i] = item
		queue[i] = item
	}
	// Start here.
	visited[start].Cost = 0
	heap.Init(&queue)

	var curr *Item
	for {
		curr = heap.Pop(&queue).(*Item)
		// We found it!
		if curr.Node == end {
			break
		}
		// Not it.  Let's check this nodes edges.
		for _, edge := range nodes[curr.Node].Edges {
			item := visited[edge.Node]
			cost := curr.Cost + edge.Cost
			// Did we find a less expensive way to get here?
			if cost < item.Cost {
				queue.update(item, cost, curr.Node)
			}
		}
		// We failed to find the path.
		if queue.Len() == 0 {
			return nil, errors.New("Failed to find path")
		}
	}
	// Gather the path.
	path := []int{}
	for curr != nil {
		path = append(path, curr.Node)
		if curr.Node == curr.Via {
			break
		}
		curr = visited[curr.Via]
	}
	// We have to reverse the path.
	n = len(path)
	for i := 0; i < n/2; i++ {
		path[i], path[n-i-1] = path[n-i-1], path[i]
	}
	return path, nil
}
