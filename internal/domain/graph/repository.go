package graph

type Repository interface {
	SaveGraph(g *Graph) error
	LoadGraph(id string) (*Graph, error)
	DeleteGraph(id string) error
	ListGraphs() ([]*Graph, error)
}
