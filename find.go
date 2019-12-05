package practice

import (
	"fmt"
	"sort"

	"github.com/golang-collections/collections/set"
)

func FindDuplicatesBrute(data []int) []int {
	size := len(data)

	dups := make([]int, 0)
	for first := 0; first < size; first++ {
		for second := first + 1; second < size; second++ {
			if data[first] == data[second] {
				dups = append(dups, data[first])
			}
		}
	}

	return dups
}

func FindDuplicatesSort(data []int) []int {
	size := len(data)
	sort.Ints(data)
	dups := make([]int, 0)
	for i := 1; i < size; i++ {
		if data[i-1] == data[i] {
			dups = append(dups, data[i])
		}
	}

	return dups
}

func FindDuplicatesSortFind(data []int) []int {
	size := len(data)
	sort.Ints(data)
	s := set.New()
	dups := make([]int, 0)
	for i := 0; i < size; i++ {
		if s.Has(data[i]) {
			dups = append(dups, data[i])
		} else {
			s.Insert(data[i])
		}
	}

	return dups
}

func FindDuplicatesCount(data []int, n int) []int {
	size := len(data)
	dups := make([]int, 0)
	prev := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		prev[i] = 0
	}
	for i := 0; i < size; i++ {
		if prev[data[i]] >= 1 {
			dups = append(dups, data[i])
		}
		prev[data[i]]++
	}

	return dups
}

func FindMaxBrute(data []int) (int, error) {
	size := len(data)
	if size == 0 {
		return 0, fmt.Errorf("No max for 0 length array")
	}
	if size == 1 {
		return data[0], nil
	}
	maxN := data[0]
	maxC := 1
	for i := 0; i < size; i++ {
		count := 1
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxC {
			maxN, maxC = data[i], count
		}
	}
	return maxN, nil
}

func FindMaxSort(data []int) (int, error) {
	size := len(data)
	if size == 0 {
		return 0, fmt.Errorf("No max for 0 length array")
	}
	if size == 1 {
		return data[0], nil
	}
	maxN := data[0]
	maxC := 1
	count := 1
	sort.Ints(data)
	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			count++
		} else {
			count = 1
		}

		if count > maxC {
			maxN, maxC = data[i], count
		}
	}
	return maxN, nil
}

func FindMaxCount(data []int, n int) (int, error) {
	size := len(data)
	if size == 0 {
		return 0, fmt.Errorf("No max for 0 length array")
	}
	if size == 1 {
		return data[0], nil
	}
	maxN := data[0]
	maxC := 1
	prev := make([]int, n+1)
	for i := 1; i < size; i++ {
		prev[data[i]]++
		if prev[data[i]] > maxC {
			maxN, maxC = data[i], prev[data[i]]
		}
	}
	return maxN, nil
}

func FindMajorityCancelation(data []int) (int, bool) {
	size := len(data)
	if size == 0 {
		return 0, false
	}
	if size == 1 {
		return data[0], true
	}
	majI := 0
	majC := 1
	for i := 1; i < size; i++ {
		if data[majI] == data[i] {
			majC++
		} else {
			majC--
		}
		if majC == 0 {
			majI, majC = i, 1
		}
	}
	majC = 0
	for i := 0; i < size; i++ {
		if data[i] == data[majI] {
			majC++
		}
	}
	if majC > size/2 {
		return data[majI], true
	}
	return 0, false
}

func FindMissingCount(data []int) (int, bool) {
	size := len(data)
	if size == 0 {
		return 0, false
	}
	if size == 1 {
		if data[0] == 1 {
			return 2, true
		} else if data[0] == 2 {
			return 1, true
		}
		return 0, false
	}
	sorted := make([]int, size+2)
	for i := 0; i < size; i++ {
		sorted[data[i]]++
	}
	size += 2
	for i := 1; i < size; i++ {
		if sorted[i] == 0 {
			return i, true
		}
	}

	return 0, false
}

func FindMissingSum(data []int) (int, bool) {
	size := len(data)
	if size == 0 {
		return 0, false
	}

	sumExpected := size + 1
	sumActual := 0
	for i := 0; i < size; i++ {
		sumExpected += i + 1
		sumActual += data[i]
	}

	return sumExpected - sumActual, true
}

func FindMissingXOR(data []int) (int, bool) {
	size := len(data)
	if size == 0 {
		return 0, false
	}

	collected := size + 1
	for i := 0; i < size; i++ {
		collected ^= i + 1
		collected ^= data[i]
	}

	return collected, true
}

func FindSumTwoSort(data []int, value int) (int, int, bool) {
	size := len(data)
	if size < 2 {
		return 0, 0, false
	}

	sort.Ints(data)
	for left, right := 0, size-1; left < right; {
		sum := data[left] + data[right]
		if sum == value {
			return data[left], data[right], true
		}
		if sum < value {
			left++
		} else {
			right--
		}
	}

	return 0, 0, false
}

func FindSumTwoMap(data []int, value int) (int, int, bool) {
	size := len(data)
	if size < 2 {
		return 0, 0, false
	}

	pairs := make(map[int]int, 0)
	for i := 0; i < size; i++ {
		if first, ok := pairs[data[i]]; ok {
			if first <= data[i] {
				return first, data[i], ok
			}
			return data[i], first, ok
		}
		pairs[value-data[i]] = data[i]
	}

	return 0, 0, false
}

func FindMinAbsSum(data []int) (int, int, bool) {
	size := len(data)
	if size < 2 {
		return 0, 0, false
	}
	if size == 2 {
		return data[0], data[1], true
	}

	sort.Ints(data)
	minLeft, minRight := 0, size-1
	minSum := absInt(data[minLeft] + data[minRight])
	for left, right := minLeft, minRight; left < right; {
		sum := data[left] + data[right]
		if sum == 0 {
			return data[left], data[right], true
		}

		abs := absInt(sum)
		if abs < minSum {
			minLeft, minRight, minSum = left, right, abs
		}

		if sum < 0 {
			left++
		} else {
			right--
		}
	}

	return data[minLeft], data[minRight], true
}

func absInt(c int) int {
	if c < 0 {
		return c * -1
	}
	return c
}

func FindBitonicMaxBSearch(data []int) (int, bool) {
	size := len(data)
	if size < 3 {
		return 0, false
	}
	left, right := 0, size-1
	mid := right / 2
	for left < right {
		if data[mid-1] <= data[mid] && data[mid+1] <= data[mid] {
			return data[mid], true
		}
		if data[mid-1] >= data[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = left + (right-left)/2
	}

	return 0, false
}
