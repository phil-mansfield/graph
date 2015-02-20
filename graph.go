package graph

import (
	"fmt"
)

// Graph is the underlying datatype used for union-find operations.
type Graph struct {
	us, vs, ids, sizes []int
	nodeCount, largestGroup int
	unionCalled bool
}

// New creates a new Graph consisting of the specified unweighted and 
// undirected edges. The ith edge is considered to be the pair (u, v) = 
// (us[i], vs[i]). It is assumed that for all edges, u, v < nodeCount.
func New(nodeCount int, us, vs []int) *Graph {
	if len(us) != len(vs) {
		panic(fmt.Sprintf("Length of us and vs must be the same, " +
			"but are (%d, %d).", len(us), len(vs)))
	}

	g := new(Graph)

	g.us = us
	g.vs = vs
	g.nodeCount = nodeCount
	g.ids = make([]int, g.nodeCount)
	g.sizes = make([]int, g.nodeCount)

	for i := 0; i < len(us); i++ {
		if us[i] >= nodeCount || vs[i] >= nodeCount {
			panic(fmt.Sprintf("Edge %d: (%d, %d) has a node larger" + 
				" than the maximum: %d", i, us[i], vs[i], nodeCount))
		}
	}

	return g
}

// NodeCount returns the number of distinct nodes in g.
func (g *Graph) NodeCount() int { return g.nodeCount }

// LargestGroup returns the ID of the largest group in g.
func (g *Graph) LargestGroup() int { return g.largestGroup }
