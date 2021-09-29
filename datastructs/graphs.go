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

// directed graphs in this case and no cycles considered.
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

// directed graphs in this case and no cycles considered.
func (g *Graph) BreadthFirst(start int) []int {
	q := Queue{}
	q.Enqueue(start)

	depthList := make([]int, 0)

	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		depthList = append(depthList, val)
		if v, ok := g.nodes[val]; ok {
			for i := 0; i < len(v.edges); i++ {
				q.Enqueue(v.edges[i])
			}
		} else {
			// no vertex with this edge found.
			break
		}
	}
	return depthList
}

// directed graphs in this case and no cycles considered.
func (g *Graph) HasPath(start int, end int) bool {
	s := Stack{}
	s.Push(start)

	for !s.IsEmpty() {
		val, _ := s.Pop()
		if val == end {
			return true
		}
		if v, ok := g.nodes[val]; ok {
			for i := 0; i < len(v.edges); i++ {
				s.Push(v.edges[i])
			}
		} else {
			// no vertex with this edge found.
			break
		}
	}
	return false
}

// directed graphs in this case and no cycles considered.
func (g *Graph) HasPathRecursive(start int, end int) bool {
	if start == end {
		return true
	}

	for _, v := range g.nodes[start].edges {
		if g.HasPathRecursive(v, end) {
			return true
		}
	}

	return false
}

// undirected graphs in this case and cycles considered.
func (g *Graph) HasPathRecursiveUndirected(start int, end int) bool {
	visited := make(map[int]bool)
	return g.undirectedpath(start, end, visited)
}

func (g *Graph) undirectedpath(start int, end int, visited map[int]bool) bool {
	if start == end {
		return true
	}
	if visited[start] {
		return false
	}
	visited[start] = true

	for _, v := range g.nodes[start].edges {
		if g.undirectedpath(v, end, visited) {
			return true
		}
	}
	return false

}

// find the number of groups within a graph. A group is a set of connected nodes.
func (g *Graph) FindGroupCount() int {
	cnt := 0
	visited := make(map[int]bool)

	// loop through each node in the graph and check its adjacency list and mark them.
	for k, _ := range g.nodes {
		if visited[k] {
			continue
		}
		// if k is not visited already, traverse through its adjacency list and mark them as visited.
		if !visited[k] {
			g.traverse(k, visited)
			cnt++
		}

	}

	return cnt
}

func (g *Graph) traverse(pos int, visited map[int]bool) {
	if visited[pos] {
		return
	}
	visited[pos] = true

	for _, v := range g.nodes[pos].edges {
		g.traverse(v, visited)
	}
}
