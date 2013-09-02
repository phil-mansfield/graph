package graph

import (
	"fmt"
)

// QueryType is a flag representing a type of question that can be asked about
// a group of connected nodes. Query flags are passed to the method 
// graph.Query
type QueryType uint8
const (
	Size QueryType = iota
	Width
	Height
)

type Graph struct {
	us, vs, ids, sizes, lefts, rights, tops, bottoms []uint32
	nodeCount, largestGroup uint32
	unionCalled bool
}

// New creates a new graph consisting of the specified unweighted and 
// undirected edges. The ith edge is considered to be the pair (u, v) = 
// (us[i], vs[i]). It is assumed that for all edges, u, v < nodeCount.
func New(nodeCount uint32, us, vs []uint32) *Graph {
	if len(us) != len(vs) {
		panic(fmt.Sprintf("Length of 'us' and 'vs' must be the same, but are" +
			" (%d, %d).", len(us), len(vs)))
	}

	g := new(Graph)

	g.us = make([]uint32, len(us))
	g.vs = make([]uint32, len(vs))
	g.nodeCount = nodeCount

	for i := 0; i < len(us); i++ {
		g.us[i] = us[i]
		g.vs[i] = vs[i]

		if us[i] >= nodeCount || vs[i] >= nodeCount {
			panic(fmt.Sprintf("Edge %d: (%d, %d) has a node larger" + 
				" than the maximum: %d", i, us[i], vs[i], nodeCount))
		}
	}

	return g
}

func (g *Graph) NodeCount() uint32 { return g.nodeCount }
func (g *Graph) LargestGroup() uint32 { return g.largestGroup }
