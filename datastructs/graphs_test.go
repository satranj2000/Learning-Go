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
