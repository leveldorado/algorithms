package algorithms

type Graph struct {
	Nodes map[Node][]Edge
}

type Node struct {
	Name string
}

type Edge struct {
	To   Node
	Cost int
}

func (g Graph) DijkstrasPathSearch(start, finish Node) ([]Node, int) {
	costs := map[Node]int{}
	parents := map[Node]Node{}
	for _, edge := range g.Nodes[start] {
		costs[edge.To] = edge.Cost
		parents[edge.To] = start
	}
	processed := map[Node]bool{}
	for {
		node, cost := g.getLowestCostNode(costs, processed)
		if node.Name == "" {
			break
		}
		for _, edge := range g.Nodes[node] {
			newCost := cost + edge.Cost
			if currentCost, ok := costs[edge.To]; ok && currentCost <= newCost {
				continue
			}
			costs[edge.To] = newCost
			parents[edge.To] = node
		}
		processed[node] = true
	}
	parent, ok := parents[finish]
	if !ok {
		return nil, 0
	}
	path := []Node{finish}
	for {
		path = append(path, parent)
		parent, ok = parents[parent]
		if !ok {
			break
		}
	}
	return path, costs[finish]
}

func (g Graph) getLowestCostNode(costs map[Node]int, proceedNodes map[Node]bool) (Node, int) {
	var lowestCostNode Node
	var lowestCost int
	for node, cost := range costs {
		if proceedNodes[node] {
			continue
		}
		if lowestCostNode.Name == "" || lowestCost > cost {
			lowestCostNode = node
			lowestCost = cost
		}
	}
	return lowestCostNode, lowestCost
}
