package search

func LinearUnsorted(data []int, value int) bool {
	size := len(data)

	for i := 0; i < size; i++ {
		if data[i] == value {
			return true
		}
	}

	return false
}

func LinearSorted(data []int, value int) bool {
	size := len(data)

	for i := 0; i < size && data[i] <= value; i++ {
		if data[i] == value {
			return true
		}
	}

	return false
}

func Binary(data []int, value int) bool {
	low := 0
	high := len(data) - 1
	var mid int

	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		}
		if value < data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}

func BinaryRecursive(data []int, value int) bool {
	return binaryRecursiveSub(data, 0, len(data)-1, value)
}

func binaryRecursiveSub(data []int, low, high, value int) bool {
	if low > high {
		return false
	}
	mid := low + (high-low)/2
	if data[mid] == value {
		return true
	}

	if value < data[mid] {
		return binaryRecursiveSub(data, low, mid-1, value)
	}

	return binaryRecursiveSub(data, mid+1, high, value)
}
