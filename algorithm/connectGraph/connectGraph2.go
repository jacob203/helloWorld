package connectGraph2

type graphValue struct {
	parentId int
	subsetCount int
}
type graph struct{
	data []graphValue
}

func NewGraph(nodeCount int) *graph  {
	g := &graph{data:make([]graphValue, nodeCount, nodeCount)}
	for i, _ := range g.data {
		g.data[i].parentId = i
		g.data[i].subsetCount = 1
	}
	return g
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
		g.data[aLabel].parentId = bLabel
		g.data[bLabel].subsetCount += g.data[aLabel].subsetCount
	}
}

func (g *graph) Query(a int, b int) bool {
	if !g.isValid(a) || !g.isValid(b) {
		return false
	}

	return g.find(a) == g.find(b)
}

func (g *graph) find(a int) int {
	curId := a
	for a != g.data[a].parentId {
		a = g.data[a].parentId
	}

	g.data[curId].parentId = a
	return a
}
