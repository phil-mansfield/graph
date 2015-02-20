package graph

import (
	"math/rand"

	"testing"
)

func makeSmallGraph() *Graph {
	us := []int{1, 1, 3, 7, 9,  8, 2}
	vs := []int{2, 3, 5, 8, 10, 9, 3}
	return New(11, us, vs)
}

func inArray(target int, xs []int) bool {
	for _, x := range xs {
		if x == target {
			return true
		}
	}
	return false
}

func TestFind(t *testing.T) {
	tests := []struct {
		target int
		validIds []int
	}{
		{0, []int{0}},
		{1, []int{1, 2, 3, 5}},
		{2, []int{1, 2, 3, 5}},
		{3, []int{1, 2, 3, 5}},
		{4, []int{4}},
		{5, []int{1, 2, 3, 5}},
		{6, []int{6}},
		{7, []int{7, 8, 9, 10}},
		{8, []int{7, 8, 9, 10}},
		{9, []int{7, 8, 9, 10}},
		{10, []int{7, 8, 9, 10}},
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

func BenchmarkUnion(b *testing.B) {
	nodes := 256 * 256 * 256
	us := make([]int, nodes * 3)
	vs := make([]int, nodes * 3)
	
	for i := range us {
		u := rand.Intn(nodes)
		v := (u + rand.Intn(15)) % nodes
		us[i], vs[i] = u, v
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		g := New(nodes, us, vs)
		g.Union()

		for j := 0; j < nodes; j++ {
			g.Find(j)
		}
	}
}
