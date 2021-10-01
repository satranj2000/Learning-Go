package datastructs

import "testing"

func TestDepthFirst(t *testing.T) {
	n1 := &GraphNode{Val: 1, edges: []int{2, 3}}
	n2 := &GraphNode{Val: 2, edges: []int{4}}
	n3 := &GraphNode{Val: 3, edges: []int{5}}
	n4 := &GraphNode{Val: 4, edges: []int{3}}
	n5 := &GraphNode{Val: 5, edges: []int{}}
	n6 := &GraphNode{Val: 6, edges: []int{5}}

	g := Graph{}

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)

	startpos := []int{1, 6, 4}
	outputs := [][]int{
		{1, 3, 5, 2, 4, 3, 5},
		{6, 5},
		{4, 3, 5},
	}

	for i := range outputs {
		res := g.DepthFirst(startpos[i])
		if !CompareArrays(res, outputs[i]) {
			t.Errorf("Expected %v, got %v", outputs[0], res)
		}
	}

}

func TestBreadthFirst(t *testing.T) {
	n1 := &GraphNode{Val: 1, edges: []int{2, 3}}
	n2 := &GraphNode{Val: 2, edges: []int{4}}
	n3 := &GraphNode{Val: 3, edges: []int{5}}
	n4 := &GraphNode{Val: 4, edges: []int{3}}
	n5 := &GraphNode{Val: 5, edges: []int{}}
	n6 := &GraphNode{Val: 6, edges: []int{5}}

	g := Graph{}

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)

	startpos := []int{1, 6, 4}
	outputs := [][]int{
		{1, 2, 3, 4, 5, 3, 5},
		{6, 5},
		{4, 3, 5},
	}

	for i := range outputs {
		res := g.BreadthFirst(startpos[i])
		if !CompareArrays(res, outputs[i]) {
			t.Errorf("Expected %v, got %v", outputs[0], res)
		}
	}

}

func TestHasPath(t *testing.T) {
	n1 := &GraphNode{Val: 1, edges: []int{2, 3}}
	n2 := &GraphNode{Val: 2, edges: []int{4}}
	n3 := &GraphNode{Val: 3, edges: []int{5}}
	n4 := &GraphNode{Val: 4, edges: []int{3}}
	n5 := &GraphNode{Val: 5, edges: []int{}}
	n6 := &GraphNode{Val: 6, edges: []int{5}}

	g := Graph{}

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)

	if !g.HasPath(1, 5) {
		t.Errorf("Expected a path but did not receive one")
	}

	if g.HasPath(6, 4) {
		t.Errorf("Did not expect a path but got one")
	}

	if !g.HasPath(2, 2) {
		t.Errorf("Expected a path but did not receive one")
	}
}

func TestHasPathRecursive(t *testing.T) {
	n1 := &GraphNode{Val: 1, edges: []int{2, 3}}
	n2 := &GraphNode{Val: 2, edges: []int{4}}
	n3 := &GraphNode{Val: 3, edges: []int{5}}
	n4 := &GraphNode{Val: 4, edges: []int{3}}
	n5 := &GraphNode{Val: 5, edges: []int{}}
	n6 := &GraphNode{Val: 6, edges: []int{5}}

	g := Graph{}

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)

	if !g.HasPathRecursive(1, 5) {
		t.Errorf("Expected a path but did not receive one")
	}

	if g.HasPathRecursive(6, 4) {
		t.Errorf("Did not expect a path but got one")
	}

	if !g.HasPathRecursive(2, 2) {
		t.Errorf("Expected a path but did not receive one")
	}
}

// there could be cycles in the graph and adjacency list is what is created here.
func TestHasPathRecursiveUndirected(t *testing.T) {
	n1 := &GraphNode{Val: 1, edges: []int{2, 3}}
	n2 := &GraphNode{Val: 2, edges: []int{1, 4}}
	n3 := &GraphNode{Val: 3, edges: []int{1, 4, 5}}
	n4 := &GraphNode{Val: 4, edges: []int{2, 3}}
	n5 := &GraphNode{Val: 5, edges: []int{3, 6}}
	n6 := &GraphNode{Val: 6, edges: []int{5}}
	n7 := &GraphNode{Val: 7, edges: []int{8}}
	n8 := &GraphNode{Val: 8, edges: []int{7}}

	g := Graph{}

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)
	g.AddNode(n7)
	g.AddNode(n8)

	if !g.HasPathRecursiveUndirected(1, 5) {
		t.Errorf("Expected a path but did not receive one")
	}

	if !g.HasPathRecursiveUndirected(6, 4) {
		t.Errorf("Expected a path but did not receive one")
	}

	if !g.HasPathRecursiveUndirected(2, 2) {
		t.Errorf("Expected a path but did not receive one")
	}

	if g.HasPathRecursiveUndirected(7, 2) {
		t.Errorf("Did not expect a path between 2 and 7 but got one")
	}
}

func TestFindGraphGroupCount1(t *testing.T) {
	n1 := &GraphNode{Val: 1, edges: []int{2, 3}}
	n2 := &GraphNode{Val: 2, edges: []int{1, 4}}
	n3 := &GraphNode{Val: 3, edges: []int{1, 4, 5}}
	n4 := &GraphNode{Val: 4, edges: []int{2, 3}}
	n5 := &GraphNode{Val: 5, edges: []int{3, 6}}
	n6 := &GraphNode{Val: 6, edges: []int{5}}
	n7 := &GraphNode{Val: 7, edges: []int{8}}
	n8 := &GraphNode{Val: 8, edges: []int{7}}

	g := Graph{}

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)
	g.AddNode(n7)
	g.AddNode(n8)

	res := g.FindGroupCount()
	if res != 2 {
		t.Errorf("Expected the graph count to be 2, but got  %v", res)
	}

}

func TestFindGraphGroupCount2(t *testing.T) {
	n1 := &GraphNode{Val: 1, edges: []int{2, 3}}
	n2 := &GraphNode{Val: 2, edges: []int{1, 4}}
	n3 := &GraphNode{Val: 3, edges: []int{1, 4, 5}}
	n4 := &GraphNode{Val: 4, edges: []int{2, 3}}
	n5 := &GraphNode{Val: 5, edges: []int{3, 6}}
	n6 := &GraphNode{Val: 6, edges: []int{5}}
	n7 := &GraphNode{Val: 7, edges: []int{8}}
	n8 := &GraphNode{Val: 8, edges: []int{7}}
	n9 := &GraphNode{Val: 9, edges: []int{}}

	g := Graph{}

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)
	g.AddNode(n7)
	g.AddNode(n8)
	g.AddNode(n9)

	res := g.FindGroupCount()
	if res != 3 {
		t.Errorf("Expected the graph count to be 3, but got  %v", res)
	}

}

func TestFindGraphGroupCount3(t *testing.T) {
	n1 := &GraphNode{Val: 1, edges: []int{2, 3}}
	n2 := &GraphNode{Val: 2, edges: []int{1, 4}}
	n3 := &GraphNode{Val: 3, edges: []int{1, 4, 5}}
	n4 := &GraphNode{Val: 4, edges: []int{2, 3}}
	n5 := &GraphNode{Val: 5, edges: []int{3, 6}}
	n6 := &GraphNode{Val: 6, edges: []int{5}}
	g := Graph{}

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)

	res := g.FindGroupCount()
	if res != 1 {
		t.Errorf("Expected the graph count to be 1, but got  %v", res)
	}

}

func TestFindMaxGraphGroupCount(t *testing.T) {
	n1 := &GraphNode{Val: 1, edges: []int{2, 3}}
	n2 := &GraphNode{Val: 2, edges: []int{1, 4}}
	n3 := &GraphNode{Val: 3, edges: []int{1, 4, 5}}
	n4 := &GraphNode{Val: 4, edges: []int{2, 3}}
	n5 := &GraphNode{Val: 5, edges: []int{3, 6}}
	n6 := &GraphNode{Val: 6, edges: []int{5}}
	g := Graph{}

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)

	res := g.FindMaxGroupNodeCount()
	if res != 6 {
		t.Errorf("Expected the max graph count to be 6, but got  %v", res)
	}

}

func TestFindMaxGraphGroupCount2(t *testing.T) {
	n1 := &GraphNode{Val: 1, edges: []int{2, 3}}
	n2 := &GraphNode{Val: 2, edges: []int{1, 4}}
	n3 := &GraphNode{Val: 3, edges: []int{1, 4, 5}}
	n4 := &GraphNode{Val: 4, edges: []int{2, 3}}
	n5 := &GraphNode{Val: 5, edges: []int{3, 6}}
	n6 := &GraphNode{Val: 6, edges: []int{5}}
	n7 := &GraphNode{Val: 7, edges: []int{8}}
	n8 := &GraphNode{Val: 8, edges: []int{7}}
	n9 := &GraphNode{Val: 9, edges: []int{}}

	g := Graph{}

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)
	g.AddNode(n7)
	g.AddNode(n8)
	g.AddNode(n9)

	res := g.FindMaxGroupNodeCount()
	if res != 6 {
		t.Errorf("Expected the max graph count to be 6, but got  %v", res)
	}

}

func TestFindMaxGraphGroupCount3(t *testing.T) {
	n1 := &GraphNode{Val: 1, edges: []int{2, 3}}
	n2 := &GraphNode{Val: 2, edges: []int{1, 4}}
	n3 := &GraphNode{Val: 3, edges: []int{1, 4}}
	n4 := &GraphNode{Val: 4, edges: []int{2, 3}}
	n7 := &GraphNode{Val: 7, edges: []int{8, 9}}
	n8 := &GraphNode{Val: 8, edges: []int{7}}
	n9 := &GraphNode{Val: 9, edges: []int{9}}

	g := Graph{}

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n7)
	g.AddNode(n8)
	g.AddNode(n9)

	res := g.FindMaxGroupNodeCount()
	if res != 4 {
		t.Errorf("Expected the max graph count to be 4, but got  %v", res)
	}

}

func TestShortestPath(t *testing.T) {
	n1 := &GraphNode{Val: 1, edges: []int{2, 3}}
	n2 := &GraphNode{Val: 2, edges: []int{1, 4}}
	n3 := &GraphNode{Val: 3, edges: []int{1, 4}}
	n4 := &GraphNode{Val: 4, edges: []int{2, 3}}
	n7 := &GraphNode{Val: 7, edges: []int{8, 9}}
	n8 := &GraphNode{Val: 8, edges: []int{7}}
	n9 := &GraphNode{Val: 9, edges: []int{9}}

	g := Graph{}

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n7)
	g.AddNode(n8)
	g.AddNode(n9)

	res := g.ShortestPath(1, 4)
	if res != 2 {
		t.Errorf("Expected the shortest path between 1 and 4 as 2. But, got %v", res)
	}

	res = g.ShortestPath(7, 9)
	if res != 1 {
		t.Errorf("Expected the shortest path between 7 and 9 as 1. But, got %v", res)
	}

	res = g.ShortestPath(1, 9)
	if res != -1 {
		t.Errorf("Expected the shortest path between 7 and 9 as -1(not found). But, got %v", res)
	}
}
