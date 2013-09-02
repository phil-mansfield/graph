package graph

import (
	"testing"
)

func makeSmallGraph() *Graph {
	us := []uint32{1, 1, 3, 7, 9,  8, 2}
	vs := []uint32{2, 3, 5, 8, 10, 9, 3}
	return New(11, us, vs)
}

func inArray(target uint32, xs []uint32) bool {
	for _, x := range xs {
		if x == target {
			return true
		}
	}
	return false
}

func TestFind(t *testing.T) {
	tests := []struct {
		target uint32
		validIds []uint32
	}{
		{0, []uint32{0}},
		{1, []uint32{1, 2, 3, 5}},
		{2, []uint32{1, 2, 3, 5}},
		{3, []uint32{1, 2, 3, 5}},
		{4, []uint32{4}},
		{5, []uint32{1, 2, 3, 5}},
		{6, []uint32{6}},
		{7, []uint32{7, 8, 9, 10}},
		{8, []uint32{7, 8, 9, 10}},
		{9, []uint32{7, 8, 9, 10}},
		{10, []uint32{7, 8, 9, 10}},
	}

	g := makeSmallGraph()
	g.Union()
	
	for _, test := range tests {
		res := g.Find(test.target)
		if !inArray(res, test.validIds) {
			t.Errorf("Node %d should have a group id in range %v, but had %d.",
				test.target, test.validIds, res)
		}

		// This is slow and quadratic, but is fine for small graphs.
		for _, n := range test.validIds {
			nRes := g.Find(n)
			if res != nRes {
				t.Errorf("Nodes %d and %d should be in same group, but are" + 
					" in groups %d and %d.", test.target, n, res, nRes)
			}
		}
	}
}

func TestLargestGroup(t *testing.T) {
	g := makeSmallGraph()
	g.Union()

	if g.Query(Size, g.LargestGroup()) != 4 {
		t.Errorf("Graph's largest group should have size 4 but has size %d.",
			g.Query(Size, g.LargestGroup()))
	}
}
