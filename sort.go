package algorithms

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	pivot := list[0]
	var left, right []int
	for i := range list {
		if i == 0 {
			continue
		}
		if list[i] < pivot {
			left = append(left, list[i])
			continue
		}
		right = append(right, list[i])
	}
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

func QuickSortMiddle(list []int) []int {
	if len(list) < 2 {
		return list
	}
	pivotI := len(list) / 2
	var left, right []int
	pivot := list[pivotI]
	for i := range list {
		if pivotI != i {
			continue
		}
		if list[i] < pivot {
			left = append(left, list[i])
			continue
		}
		right = append(right, list[i])
	}
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
