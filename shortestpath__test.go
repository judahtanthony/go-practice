package practice

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	input := [][]Edge{
		[]Edge{Edge{1, 7}, Edge{2, 2}, Edge{3, 3}},
		[]Edge{Edge{0, 3}, Edge{2, 3}, Edge{4, 4}},
		[]Edge{Edge{0, 2}, Edge{1, 3}, Edge{4, 4}, Edge{7, 1}},
		[]Edge{Edge{0, 3}, Edge{11, 2}},
		[]Edge{Edge{1, 4}, Edge{2, 4}, Edge{5, 5}},
		[]Edge{Edge{4, 5}, Edge{7, 3}},
		[]Edge{Edge{7, 2}, Edge{12, 2}},
		[]Edge{Edge{2, 1}, Edge{5, 3}, Edge{6, 2}},
		[]Edge{Edge{11, 4}, Edge{9, 6}, Edge{10, 4}},
		[]Edge{Edge{11, 4}, Edge{8, 6}, Edge{10, 4}},
		[]Edge{Edge{8, 4}, Edge{9, 4}, Edge{12, 5}},
		[]Edge{Edge{3, 2}, Edge{8, 4}, Edge{9, 4}},
		[]Edge{Edge{6, 2}, Edge{10, 5}},
	}
	expected := []int{0, 2, 7, 6, 12}
	actual, err := Dijkstra(input, 0, 12)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v\nActual: %v\n", expected, actual)
	}
}

func TestAStar(t *testing.T) {
	input := []Node{
		Node{10, []Edge{Edge{1, 7}, Edge{2, 2}, Edge{3, 3}}},
		Node{9, []Edge{Edge{0, 3}, Edge{2, 3}, Edge{4, 4}}},
		Node{7, []Edge{Edge{0, 2}, Edge{1, 3}, Edge{4, 4}, Edge{7, 1}}},
		Node{8, []Edge{Edge{0, 3}, Edge{11, 2}}},
		Node{8, []Edge{Edge{1, 4}, Edge{2, 4}, Edge{5, 5}}},
		Node{6, []Edge{Edge{4, 5}, Edge{7, 3}}},
		Node{3, []Edge{Edge{7, 2}, Edge{12, 2}}},
		Node{6, []Edge{Edge{2, 1}, Edge{5, 3}, Edge{6, 2}}},
		Node{4, []Edge{Edge{11, 4}, Edge{9, 6}, Edge{10, 4}}},
		Node{4, []Edge{Edge{11, 4}, Edge{8, 6}, Edge{10, 4}}},
		Node{3, []Edge{Edge{8, 4}, Edge{9, 4}, Edge{12, 5}}},
		Node{6, []Edge{Edge{3, 2}, Edge{8, 4}, Edge{9, 4}}},
		Node{0, []Edge{Edge{6, 2}, Edge{10, 5}}},
	}
	expected := []int{0, 2, 7, 6, 12}
	actual, err := AStar(input, 0, 12)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v\nActual: %v\n", expected, actual)
	}
}
