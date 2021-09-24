package datastructs

//defines one vertex and its edges/adjacency list or nodes that are linked to this node
type GraphNode struct {
	Val   int
	edges []int
}

// group of Graph Nodes make a graph
type Graph struct {
	nodes map[int]*GraphNode
}

func (g *Graph) Count() int {
	return len(g.nodes)
}

func (g *Graph) AddNode(graph *GraphNode) {
	if g.nodes == nil {
		g.nodes = make(map[int]*GraphNode)
	}
	g.nodes[graph.Val] = graph
}

func (gnode *GraphNode) AddEdges(e int) bool {
	gnode.edges = append(gnode.edges, e)
	return true
}

func (g *Graph) DepthFirst(start int) []int {
	s := Stack{}
	s.Push(start)

	depthList := make([]int, 0)

	for !s.IsEmpty() {
		val, _ := s.Pop()
		depthList = append(depthList, val)
		if v, ok := g.nodes[val]; ok {
			for i := 0; i < len(v.edges); i++ {
				s.Push(v.edges[i])
			}
		} else {
			// no vertex with this edge found.
			break
		}
	}

	return depthList
}
