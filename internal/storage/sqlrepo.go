package storage

import (
	"database/sql"
	"strconv"

	"github.com/gola/internal/conf"
	"github.com/gola/internal/model"
	_ "github.com/lib/pq"
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
	rslt := make([]model.Node, 0)

	rows, err := r.db.Query("SELECT id, name, node_type FROM node")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := model.Node{}
		if errz := rows.Scan(&n.ID, &n.Name, &n.NodeType); errz != nil {
			return nil, errz
		}
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

func (r SQLRepo) CreateEdge(e *model.Edge) error {
	if _, err := r.db.Exec(
		"INSERT INTO edge (left_id, right_id) VALUES ($1, $2)",
		e.LeftID,
		e.RightID,
	); err != nil {
		return err
	}
	return nil
}

func (r SQLRepo) GetEdges() ([]model.Edge, error) {
	rslt := make([]model.Edge, 1)

	rows, err := r.db.Query("SELECT left_id, right_id FROM edge")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		e := model.Edge{}
		if errz := rows.Scan(&e.LeftID, &e.RightID); errz != nil {
			return nil, errz
		}
		rslt = append(rslt, e)
	}

	return rslt, nil
}

func (r SQLRepo) GetInputEdgesByNodeID(node_id int) ([]model.Edge, error) {
	rslt := make([]model.Edge, 0)

	rows, err := r.db.Query("SELECT left_id, right_id FROM edge WHERE right_id=$1", node_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		e := model.Edge{}
		if errz := rows.Scan(&e.LeftID, &e.RightID); errz != nil {
			return nil, errz
		}
		rslt = append(rslt, e)
	}

	return rslt, nil
}

func (r SQLRepo) GetOutputEdgesByNodeID(node_id int) ([]model.Edge, error) {
	rslt := make([]model.Edge, 0)

	rows, err := r.db.Query("SELECT left_id, right_id FROM edge WHERE left_id=$1", node_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		e := model.Edge{}
		if errz := rows.Scan(&e.LeftID, &e.RightID); errz != nil {
			return nil, errz
		}
		rslt = append(rslt, e)
	}

	return rslt, nil
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
