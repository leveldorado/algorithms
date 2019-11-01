package algorithms

func BinarySearch(list []int, n int) int {
	high := len(list) - 1
	low := 0
	for low <= high {
		mid := low + (high-low)/2
		el := list[mid]
		if el == n {
			return mid
		}
		if el > n {
			high = mid - 1
			continue
		}
		low = mid + 1
	}
	return -1
}
