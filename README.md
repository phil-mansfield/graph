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
graphs's `Union()` method to link nodes together based off of those edges.

After grouping the graph's nodes, it is possible to find the IDs of each node
by calling the graph's `Find(node)` method. It is also possible to find
properties of a given group of nodes via the graph's `Query(query, groupID)`
method.

Implementation Notes
--------------------


Runtime
-------

Calls to `Union` execute in at most `O(N log*(N))` and calls to `Find`
execute in at most `O(log*(N))`, where `N` is the minimum of the edge count
and the node count, and `log*(N)` is the number of times one can take the
logarithm of `N` before reaching 1 or less. `log*(N)` is less than 5 for any
`N` with less than 20,000 digits, so it is safe to treat these as constant
factors for our purposes.