package graph

type Node struct {
	ID    string
	Label string
	Data  map[string]any
}

type Edge struct {
	From string
	To   string
	Type string
}

type Graph struct {
	ID      string
	Nodes   map[string]*Node
	Edges   map[string][]*Edge // adjacency list
}

func NewGraph(id string) *Graph {
	return &Graph{
		ID:    id,
		Nodes: make(map[string]*Node),
		Edges: make(map[string][]*Edge),
	}
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes[n.ID] = n
}

func (g *Graph) AddEdge(from, to, edgeType string) {
	g.Edges[from] = append(g.Edges[from], &Edge{From: from, To: to, Type: edgeType})
}

func (g *Graph) Neighbors(id string) []*Node {
	var res []*Node
	for _, e := range g.Edges[id] {
		if n, ok := g.Nodes[e.To]; ok {
			res = append(res, n)
		}
	}
	return res
}
