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

func (g *Graph) FindMaxGroupNodeCount() int {
	max := 0
	visited := make(map[int]bool)

	// loop through each node in the graph and check its adjacency list and mark them.
	for k, _ := range g.nodes {
		if visited[k] {
			continue
		}
		// if k is not visited already, traverse through its adjacency list and mark them as visited.
		if !visited[k] {
			cnt := g.traverseCount(k, visited)
			if cnt > max {
				max = cnt
			}
		}

	}
	return max
}

func (g *Graph) traverseCount(pos int, visited map[int]bool) int {
	if visited[pos] {
		return 0
	}
	visited[pos] = true
	cnt := 1 // since we visited this position now set it to 1.
	for _, v := range g.nodes[pos].edges {
		cnt += g.traverseCount(v, visited) // combine all the counts
	}
	return cnt
}

func (g *Graph) ShortestPath(start int, end int) int {
	// use breadth first for finding the shortest path.
	visited := make(map[int]bool)
	// 2 queues. one to track the node we are visiting and another to track the distance of the node from start.
	q := Queue{}
	qlen := Queue{}

	// first initialized to 0 length.
	q.Enqueue(start)
	qlen.Enqueue(0)
	visited[start] = true

	for !q.IsEmpty() {
		outval, _ := q.Dequeue()
		outvalLen, _ := qlen.Dequeue()
		if outval == end {
			return outvalLen
		}

		for _, e := range g.nodes[outval].edges {
			if !visited[e] {
				q.Enqueue(e)
				qlen.Enqueue(outvalLen + 1)
				visited[e] = true
				if e == end {
					return outvalLen + 1
				}
			}
		}
	}
	return -1
}
