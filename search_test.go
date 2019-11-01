package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	list := []int{2, 6, 9, 34, 88, 130, 700, 701}
	for n, i := range map[int]int{2: 0, 3: -1, 88: 4, 701: 7} {
		require.Equal(t, i, BinarySearch(list, n))
	}
}
