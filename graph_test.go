package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGraph_DijkstrasPathSearch(t *testing.T) {
	start, finish := Node{Name: "start"}, Node{Name: "finish"}
	a, b := Node{Name: "A"}, Node{Name: "B"}
	g := Graph{Nodes: map[Node][]Edge{
		start: {
			{
				To:   a,
				Cost: 6,
			},
			{
				To:   b,
				Cost: 2,
			},
		},
		a: {
			{
				To:   finish,
				Cost: 1,
			},
		},
		b: {
			{
				To:   a,
				Cost: 3,
			},
			{
				To:   finish,
				Cost: 5,
			},
		},
		finish: {},
	}}
	path, cost := g.DijkstrasPathSearch(start, finish)
	require.Equal(t, []Node{finish, a, b, start}, path)
	require.Equal(t, 6, cost)
}
