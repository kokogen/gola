package storage

import "github.com/gola/internal/model"

type Repo interface {
	CreateNode(*model.Node) error
	GetNodes() ([]model.Node, error)
	GetNode(int) (*model.Node, error)
}
