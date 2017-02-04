# Connecting Graph
Given n nodes in a graph labeled from 1 to n. There is no edges in the graph at beginning.

You need to support the following method:
1. connect(a, b), add an edge to connect node a and node b. 
2. query(a, b)`, check if two nodes are connected

Example:
```
5 // n = 5
query(1, 2) return false
connect(1, 2)
query(1, 3) return false
connect(2, 4)
query(1, 4) return true
```

it is an issue of unionizing nodes and checking if the two nodes are connected.  
keep in mind that there are only actions to connect two nodes, but no actions to disconnect two nodes.
if there are some actions to disconnect two node groups, the union find algorithm is not fit here.
so it is fit to use union find.
1. firstly each node is a node group individually, and its group id is itself. 
2. connect two nodes, find the two nodes' group id, and put one group id to the other.
3. check if two nodes are connected, find the two nodes' group id, and check if they are the same.

how to find the group id is how to find the group node whose group id is itself.  
the connect action time complexity is n, so if there are k1 connect actions, the complexity is k1n.  
the query action time complexity is also n, so if there are k2 query actions, the complexity is k2n.  
so the final time complexity is kn, k is the max of k1 and k2.  

there is a method to optimize the complexity, which is compression path. during the union find, update the node's group node
to the root node. 
the union find's time complexity can be optimized to lgn on average, so the time complexity is klgn, but its worst time complexity is also kn.

# connectGraph2
Given n nodes in a graph labeled from 1 to n. There is no edges in the graph at beginning.

You need to support the following method:
1. connect(a, b), an edge to connect node a and node b
2. query(a), Returns the number of connected component nodes which include node a.  

example:
```
5 // n = 5
query(1) return 1
connect(1, 2)
query(1) return 2
connect(2, 4)
query(1) return 3
connect(1, 4)
query(1) return 3
```
it is also an union-find problem, but the query returns the group element count, so just do a little modification to the previous algorithm.
1. normal method: add another element to the node value to count the subset nodes
2. compression method: during find, update the node's parent node to the group node

# connectGraph3
Given n nodes in a graph labeled from 1 to n. There is no edges in the graph at beginning.

You need to support the following method:
1. connect(a, b), an edge to connect node a and node b
2. query(), Returns the number of connected component in the graph  

example:
```
5 // n = 5
query() return 5
connect(1, 2)
query() return 4
connect(2, 4)
query() return 3
connect(1, 4)
query() return 3
```
it is also an union-find problem, this time, it gets the group count.
use a group id list to record the current group ids, and when querying, the group count is the group id list size.