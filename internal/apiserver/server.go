package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gola/internal/model"
	"github.com/gola/internal/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router mux.Router
	logger logrus.Logger
	repo   storage.Repo
}

func NewServer(r *storage.Repo) *Server {
	s := &Server{
		router: *mux.NewRouter(),
		logger: *logrus.New(),
		repo:   *r,
	}

	s.configure()

	return s
}

func (s *Server) configure() {
	// ...
	s.router.HandleFunc("/nodes", s.createHandlerCreateNode()).Methods("POST")
	s.router.HandleFunc("/nodes", s.createHandlerGetAllNodes()).Methods("GET")
	s.router.HandleFunc("/nodes/{id}", s.createHandlerGetNode()).Methods("GET")
}

func (s *Server) createHandlerCreateNode() http.HandlerFunc {
	type request struct {
		Name     string `json:"name"`
		NodeType string `json:"node_type"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		re := request{}

		if err := json.NewDecoder(r.Body).Decode(&re); err != nil {
			s.logger.Error(err)
			return
		}

		node := model.Node{Name: re.Name, NodeType: re.NodeType}

		if err := s.repo.CreateNode(&node); err != nil {
			s.logger.Error(err)
			return
		}

	}
}

func (s *Server) createHandlerGetAllNodes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rslt, err := s.repo.GetNodes()

		if err != nil {
			s.logger.Error(err)
			return
		}

		if rslt != nil {
			json.NewEncoder(w).Encode(rslt)
		}
	}
}

func (s *Server) createHandlerGetNode() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			s.logger.Error(err)
			return
		}

		rslt, err := s.repo.GetNode(id)

		if err != nil {
			s.logger.Error(err)
			return
		}

		if rslt != nil {
			json.NewEncoder(w).Encode(rslt)
		}
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
