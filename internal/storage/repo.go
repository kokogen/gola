package storage

import "github.com/gola/internal/model"

type Repo interface {
	CreateNode(*model.Node) error
	GetNodes() ([]model.Node, error)
	GetNode(int) (*model.Node, error)

	CreateEdge(*model.Edge) error
	GetEdges() ([]model.Edge, error)
	GetInputEdgesByNodeID(int) ([]model.Edge, error)
	GetOutputEdgesByNodeID(int) ([]model.Edge, error)
}
