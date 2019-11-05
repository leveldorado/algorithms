package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuickSort(t *testing.T) {
	a := []int{3, 5, 12, 1, 8, 16}
	require.Equal(t, []int{1, 3, 5, 8, 12, 16}, QuickSort(a))
}
