package practice

// FibTD is an example of a top-down algorithm.
func FibTD(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return FibTD(n-1) + FibTD(n-2)
	}
}

// FibMem is an example of top-down but with an addition
// of memoization.
func FibMem(n int) int {
	fibcache := make(map[int]int, n)
	return fstep(n, fibcache)
}
func fstep(n int, cache map[int]int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		if answer, ok := cache[n]; ok {
			return answer
		}
		answer := fstep(n-1, cache) + fstep(n-2, cache)
		cache[n] = answer
		return answer
	}
}

// FibBU is an example of bottom-up tabulation dynamic
// programing.
func FibBU(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		fibcache := make([]int, n+1)
		fibcache[1] = 1
		for i := 2; i <= n; i++ {
			fibcache[i] = fibcache[i-1] + fibcache[i-2]
		}
		return fibcache[n]
	}
}
