package connectGraph

type graph struct{
	data []int
}

func NewGraph(nodeCount int) *graph  {
	g := &graph{data:make([]int, nodeCount, nodeCount)}
	for i, _ := range g.data {
		g.data[i] = i
	}
	return &graph{data:make([]int, nodeCount, nodeCount)}
}

func (g *graph) isValid(a int) bool {
	return a > 0 && a < len(g.data)
}
func (g *graph) Connect(a int, b int) {
	if !g.isValid(a) || !g.isValid(b) {
		return
	}

	aLabel := g.find(a)
	bLabel := g.find(b)
	if aLabel != bLabel {
		g.data[aLabel] = bLabel
	}
}

func (g *graph) Query(a int, b int) bool {
	if !g.isValid(a) || !g.isValid(b) {
		return false
	}

	return g.find(a) == g.find(b)
}

func (g *graph) find(a int) int {
	for a != g.data[a] {
		a = g.data[a]
	}

	return a
}

func (g *graph) update(a int, label int)  {
	for a != g.data[a] {
		cur := a
		a = g.data[a]
		g.data[cur] = label
	}
}


