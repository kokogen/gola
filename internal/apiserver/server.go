package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gola/internal/model"
	"github.com/gola/internal/storage"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	addr   string
	logger logrus.Logger
	repo   storage.Repo
}

func NewServer(r *storage.Repo, addr string) *APIServer {
	return &APIServer{
		addr:   addr,
		logger: *logrus.New(),
		repo:   *r,
	}
}

func (s *APIServer) Run() error {

	router := http.NewServeMux()

	server := http.Server{
		Addr:    s.addr,
		Handler: MiddlewareChain(MakeTrackingMiddlware, MakeLoggingMiddleware)(router),
	}

	router.HandleFunc("POST /nodes", s.handlerCreateNode())
	router.HandleFunc("GET /nodes", s.handlerGetAllNodes())
	router.HandleFunc("GET /nodes/{id}", s.handlerGetNodeById())

	router.HandleFunc("POST /edges", s.handlerCreateEdge())
	router.HandleFunc("GET /edges", s.handlerGetAllEdges())
	router.HandleFunc("GET /edges/{id}", s.handlerGetEdgesByNodeID())

	return server.ListenAndServe()
}

func (s *APIServer) handlerCreateNode() http.HandlerFunc {
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

func (s *APIServer) handlerGetAllNodes() http.HandlerFunc {
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

func (s *APIServer) handlerGetNodeById() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.PathValue("id"))
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

func (s *APIServer) handlerCreateEdge() http.HandlerFunc {
	type request struct {
		LeftID  int `json:"left_id"`
		RightID int `json:"right_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		re := request{}

		if err := json.NewDecoder(r.Body).Decode(&re); err != nil {
			s.logger.Error(err)
			return
		}

		edge := model.Edge{LeftID: re.LeftID, RightID: re.RightID}

		if err := s.repo.CreateEdge(&edge); err != nil {
			s.logger.Error(err)
			return
		}

	}
}

func (s *APIServer) handlerGetAllEdges() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rslt, err := s.repo.GetEdges()

		if err != nil {
			s.logger.Error(err)
			return
		}

		if rslt != nil {
			json.NewEncoder(w).Encode(rslt)
		}
	}
}

func (s *APIServer) handlerGetEdgesByNodeID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			s.logger.Error(err)
			return
		}

		rslt, err := s.repo.GetInputEdgesByNodeID(id)

		if err != nil {
			s.logger.Error(err)
			return
		}

		rslt1, err := s.repo.GetOutputEdgesByNodeID(id)

		if err != nil {
			s.logger.Error(err)
			return
		}

		rslt = append(rslt, rslt1...)

		if rslt != nil {
			json.NewEncoder(w).Encode(rslt)
		}
	}
}

// func (s *APIServer) handlerGetEdgesByNodeID() http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		left_id_str := r.URL.Query().Get("left_id")
// 		right_id_str := r.URL.Query().Get("right_id")

// 		switch {
// 		case (left_id_str != "") && (right_id_str == ""):
// 			left_id, err := strconv.Atoi(left_id_str)
// 			if err != nil {
// 				s.logger.Error(err)
// 				return
// 			}

// 			rslt, err := s.repo.GetInputEdgesByNodeID(left_id)

// 			if err != nil {
// 				s.logger.Error(err)
// 				return
// 			}

// 			if rslt != nil {
// 				json.NewEncoder(w).Encode(rslt)
// 			}

// 		case (left_id_str == "") && (right_id_str != ""):
// 			right_id, err := strconv.Atoi(right_id_str)
// 			if err != nil {
// 				s.logger.Error(err)
// 				return
// 			}

// 			rslt, err := s.repo.GetOutputEdgesByNodeID(right_id)

// 			if err != nil {
// 				s.logger.Error(err)
// 				return
// 			}

// 			if rslt != nil {
// 				json.NewEncoder(w).Encode(rslt)
// 			}
// 		}

// 	}
// }
