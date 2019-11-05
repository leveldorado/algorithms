package algorithms

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuickSort(t *testing.T) {
	a := []int{3, 5, 12, 1, 8, 16}
	require.Equal(t, []int{1, 3, 5, 8, 12, 16}, QuickSort(a))
}

func TestQuickSortMiddle(t *testing.T) {
	a := []int{3, 5, 12, 1, 8, 16}
	require.Equal(t, []int{1, 3, 5, 8, 12, 16}, QuickSortMiddle(a))
}

func BenchmarkQuickSort1000(b *testing.B) {
	list := generateRandomNumbers(1000)
	for n := 0; n < b.N; n++ {
		QuickSort(list)
	}
}

func BenchmarkQuickSortMiddle1000(b *testing.B) {
	list := generateRandomNumbers(1000)
	for n := 0; n < b.N; n++ {
		QuickSortMiddle(list)
	}
}

func generateRandomNumbers(size int) []int {
	var list []int
	for i := 0; i < size; i++ {
		list = append(list, rand.Intn(100))
	}
	return list
}
