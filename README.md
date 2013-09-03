graph
=====

`graph` is a small optimized union-find graph library. It is designed for use
in monte-carlo simulations. Additional documentation can be found at 
<http://godoc.org/github.com/phil-mansfield/graph>.

Functionality
-------------

The package `graph` provides a way to group connected nodes in a graph and
then poll those groups for information. To use this library, first produce
a list of graph edges (represented as a pair of integer arrays) and pass
to the funciton `graph.New(nodeCount, us, vs)`. Next, call the returned
graphs's `Union()` method to link nodes together.

After grouping the graph's nodes, it is possible to find the group IDs of any
node by calling the graph's `Find(node)` method. It is also possible to find
properties of a given group of nodes via the graph's `Query(query, groupID)`
method.

Implementation Notes
--------------------

`graph` implements union-find using a Weighted Lazy-Union with Path 
Compression (WLUPC for short). WLUPC is chosen over, for example, flood-fill
due to increased speed on "low"-edge graphs, a lighter memory footprint, the 
potential for better cache locality, and the simplicity of implementation.

At the beginning of the call to Union each node is assigned its own unique
group ID. The graph's edges are then iterated over. Whenever an edge between
two different groups is found, the ID of the smaller group is changed to the
ID of the larger group.

The group ID of each node is evaluated "lazily," meaning that it is not fully
calculated until requested. To change the ID of an entire group only the ID
of the node corresponding to the group's old ID is changed. This creates an
effective "tree" of groups. When the group ID of nodes not at the root of this
tree is queried (whether internally or through `Find`) the tree is traversed
and the group ID of the root is used.

To keep the group ID trees small and to prevent time being wasted on
tree-traversal, path-compession (ie. memoization) is used. Whenever the nodes
of a group ID tree are traversed to the root, they are then traversed a second
time and have their group ID set to the root value.

(see, eg,  <http://www.cs.princeton.edu/~rs/AlgsDS07/01UnionFind.pdf> for a
more in-depth explanation of this algorithm.)

Runtime
-------

Calls to `Union` execute in at most `O(N log*(E))` and calls to `Find`
execute in at most `O(log*(E))`, where `E` is the edge count,
and `log*(N)` is the number of times one can take the
logarithm of `N` before reaching 1 or less. `log*(N)` is less than 5 for any
`N` with less than 20,000 digits, so it is safe to treat these as constant
factors for all sane purposes.

Not all aspects are optimized as fully as they could be, but since this 
library is intended for monte carlo simulations (meaning that there has to
be at least one call to a random number generator per edge in the graph),
hyper-optimizing these routines is of limited utility. In its current form
`Union` will run in less than half the time it takes to generate every edge.