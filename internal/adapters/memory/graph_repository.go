package memory

import (
	"errors"
	"graph-football/internal/domain/graph"
)

type InMemoryGraphRepo struct {
	store map[string]*graph.Graph
}

func NewInMemoryGraphRepo() *InMemoryGraphRepo {
	return &InMemoryGraphRepo{store: make(map[string]*graph.Graph)}
}

func (r *InMemoryGraphRepo) SaveGraph(g *graph.Graph) error {
	r.store[g.ID] = g
	return nil
}

func (r *InMemoryGraphRepo) LoadGraph(id string) (*graph.Graph, error) {
	if g, ok := r.store[id]; ok {
		return g, nil
	}
	return nil, errors.New("graph not found")
}

func (r *InMemoryGraphRepo) DeleteGraph(id string) error {
	delete(r.store, id)
	return nil
}

func (r *InMemoryGraphRepo) ListGraphs() ([]*graph.Graph, error) {
	res := make([]*graph.Graph, 0, len(r.store))
	for _, g := range r.store {
		res = append(res, g)
	}
	return res, nil
}
