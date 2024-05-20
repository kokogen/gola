package storage

import (
	"database/sql"
	"strconv"

	"github.com/gola/internal/conf"
	"github.com/gola/internal/model"
)

type SQLRepo struct {
	db *sql.DB
}

func (r SQLRepo) CreateNode(n *model.Node) error {
	if err := r.db.QueryRow(
		"INSERT INTO node (name, node_type) VALUES ($1, $2) RETURNING id",
		n.Name,
		n.NodeType,
	).Scan(&n.ID); err != nil {
		return err
	}
	return nil
}

func (r SQLRepo) GetNodes() ([]model.Node, error) {
	rslt := make([]model.Node, 1)

	rows, err := r.db.Query("SELECT id, name, node_type FROM node")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := model.Node{}
		rows.Scan(&n.ID, &n.Name, &n.NodeType)
		rslt = append(rslt, n)
	}

	return rslt, nil
}

func (r SQLRepo) GetNode(id int) (*model.Node, error) {
	n := model.Node{}

	err := r.db.QueryRow("SELECT id, name, node_type FROM node WHERE id = ?", strconv.Itoa(id)).Scan(&n.ID, &n.Name, &n.NodeType)
	if err != nil {
		return nil, err
	}

	return &n, nil
}

func New(config *conf.Config) (Repo, error) {
	db, err := sql.Open("postgres", config.DbUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return SQLRepo{db: db}, nil
}
