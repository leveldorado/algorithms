package algorithms

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	pivot := list[0]
	var left, right []int
	for _, el := range list[1:] {
		if el < pivot {
			left = append(left, el)
			continue
		}
		right = append(right, el)
	}
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
