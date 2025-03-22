package boilerplates

// Disjoint Set Union (DSU) -> Union Find
// https://leetcode.com/discuss/general-discussion/1072418/Disjoint-Set-Union-(DSU)Union-Find-A-Complete-Guide
// https://leetcode.com/explore/learn/card/graph/618/disjoint-set/

type DSU struct {
	NbNodes int
	Rank    []int
	Parent  []int
}

func NewDSU(nbNodes int) *DSU {
	dsu := DSU{
		NbNodes: nbNodes,
		Rank:    make([]int, nbNodes),
		Parent:  make([]int, nbNodes),
	}
	for i := range dsu.Rank {
		dsu.Rank[i] = 1
	}
	// set itself as parent
	for i := range dsu.Parent {
		dsu.Parent[i] = i
	}

	return &dsu
}

// find the ultimate parent of the node, while compressing the path
func (dsu *DSU) Find(node int) int {
	if dsu.Parent[node] == node {
		return node
	}
	dsu.Parent[node] = dsu.Find(dsu.Parent[node])
	return dsu.Parent[node]
}

// return true if u and v belong to different sets
func (dsu *DSU) Union(u, v int) bool {
	u = dsu.Find(u)
	v = dsu.Find(v)

	if u == v {
		return false
	}

	if dsu.Rank[u] > dsu.Rank[v] {
		dsu.Parent[v] = u
		dsu.Rank[u] += dsu.Rank[v]
	} else {
		dsu.Parent[u] = v
		dsu.Rank[v] += dsu.Rank[u]
	}
	return true
}

// track the number of edges in each component
type DSUWithEdges struct {
	DSU
	Edges []int
}

func NewDSUWithEdges(nbNodes int) *DSUWithEdges {
	dsu := DSUWithEdges{
		DSU:   *NewDSU(nbNodes),
		Edges: make([]int, nbNodes),
	}
	return &dsu
}

func (dsu *DSUWithEdges) Union(u, v int) bool {
	u = dsu.Find(u)
	v = dsu.Find(v)

	// count 1 for this new edge between u and v
	dsu.Edges[u]++

	if u == v {
		return false
	}

	if dsu.Rank[u] > dsu.Rank[v] {
		dsu.Parent[v] = u
		dsu.Edges[u] += dsu.Edges[v]
		dsu.Rank[u] += dsu.Rank[v]
	} else {
		dsu.Parent[u] = v
		dsu.Edges[v] += dsu.Edges[u]
		dsu.Rank[v] += dsu.Rank[u]
	}
	return true
}
