package practice

/**
 * [1 3 5] [1 3 9] [9 5]
 * [1 3 9 5]
 */

func MergeArrays(args ...[]int) []int {
	n := len(args)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return args[0][:]
	}
	ret := make([]int, 0)
	dependencies := make(map[int][]int, 0)
	visited := make(map[int]bool)

	for i := 0; i < n; i++ {
		arr := args[i]
		for j := 0; j < len(arr); j++ {
			num := arr[j]
			if _, ok := dependencies[num]; !ok {
				dependencies[num] = make([]int, 0)
			}
			if j > 0 {
				dependencies[num] = append(dependencies[num], arr[j-1])
			}
		}
	}

	for from := range dependencies {
		TopoSort(from, dependencies, visited, &ret)
	}

	return ret
}

func TopoSort(pos int, dependencies map[int][]int, visited map[int]bool, sorted *[]int) {
	if _, ok := visited[pos]; ok {
		return
	}
	visited[pos] = true
	for _, next := range dependencies[pos] {
		TopoSort(next, dependencies, visited, sorted)
	}

	*sorted = append(*sorted, pos)
}
