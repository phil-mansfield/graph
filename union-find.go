package graph

import (
	"fmt"
)

// root returns the ID of the group at the specified node. After traversing
// down a branch of a tree it traverses back up so that all nodes along the
// path "remember" what the root was.
func (g *Graph) root(node int) int {
	groupId := node
	for g.ids[groupId] != groupId {
		groupId = g.ids[groupId]
	}
	
	for g.ids[node] != node {
		oldId := g.ids[node]
		g.ids[node] = groupId
		node = oldId
	}

	return groupId
}

// Union puts all nodes connected by edges to groups and assigns a group ID
// number to each of them. It also computes the properties of these groups so
// that Query can access them in the future.
//
// IMPLEMENTATION NOTE: This call currently runs in O(N logStar(N))
func (g *Graph) Union() {
	if g.unionCalled {
		return
	}

	for i := 0; i < len(g.ids); i++ {
		g.ids[i] = i
		g.sizes[i] = 1
	}

	for i := 0; i < len(g.us); i++ {
		u, v := g.us[i], g.vs[i]
		uRoot := g.root(u)
		vRoot := g.root(v)

		if uRoot == vRoot {
			continue
		}

		var rootId, leafId int
		if g.sizes[uRoot] > g.sizes[vRoot] {
			rootId = uRoot
			leafId = vRoot
		} else {
			rootId = vRoot
			leafId = uRoot
		}
		
		g.sizes[rootId] = g.sizes[rootId] + g.sizes[leafId]
		g.ids[leafId] = rootId

		if g.sizes[rootId] > g.sizes[g.largestGroup] {
			g.largestGroup = rootId
		}
	}

	g.unionCalled = true
}

// Find returns the ID of the group that the specified node is a member of.
// Union must be called before Find. Failing to do so will result in an error.
//
// IMPLEMENTATION NOTE: This call currently runs in O(log*(N)).
func (g *Graph) Find(node int) int {
	if !g.unionCalled {
		panic("You must call Union before Find.")
	} else if node >= g.nodeCount {
		panic(fmt.Sprintf("The node %d is larger than the accepted " +
			"nodeCount.", g.nodeCount))
	}

	return g.root(node)
}

func (g *Graph) Roots() []int {
	// TODO: Appending may be too slow. Benchmaek.
	roots := []int{}
	for i := 0; i < len(g.ids); i++ {
		if g.ids[i] == i {
			roots = append(roots, i)
		}
	}

	return roots
}
