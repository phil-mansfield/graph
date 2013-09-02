package graph

import (
	"fmt"
)

func (g *Graph) root(node uint32) uint32 {
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

func (g *Graph) Union() {
	if g.unionCalled {
		return
	}

	g.ids = make([]uint32, g.nodeCount)
	g.sizes = make([]uint32, g.nodeCount)

	for i := 0; i < len(g.ids); i++ {
		g.ids[i] = uint32(i)
		g.sizes[i] = 1
	}

	for i := 0; i < len(g.us); i++ {
		u, v := g.us[i], g.vs[i]
		uRoot := g.root(u)
		vRoot := g.root(v)

		if uRoot == vRoot {
			continue
		}

		var rootId, leafId uint32
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

func (g *Graph) Find(node uint32) uint32 {
	if !g.unionCalled {
		panic("You must call Union before Find.")
	} else if node >= g.nodeCount {
		panic(fmt.Sprintf("The node %d is larger than the accepted nodeCount.",
			g.nodeCount))
	}

	return g.root(node)
}

func (g *Graph) Query(qt QueryType, groupId uint32) int {
	if !g.unionCalled {
		panic("You must call Union before Find.")
	} else if groupId >= g.nodeCount {
		panic(fmt.Sprintf("The node %d is larger than the accepted nodeCount.",
			g.nodeCount))
	}

	switch(qt) {
	case Size:
		return int(g.sizes[groupId])
	case Height:
		panic("Not yet implemented.")
	case Width:
		panic("Not yet implemented.")
	}

	panic("Unrecognized QueryType.")
}
