package server

import (
	"SportsTestWork/config"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

type Server struct {
	context context.Context
	config  *config.Config
	db      *sql.DB
}

func Init(ctx context.Context, config *config.Config, db *sql.DB) *Server {
	s := &Server{
		context: ctx,
		config:  config,
		db:      db,
	}
	s.routes()
	return s
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			// TODO
		}
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, err error, status int) {
	w.Header().Add("Content-Type", "applications/json")
	w.WriteHeader(status)
	if err != nil {
		err := json.NewEncoder(w).Encode(e(err))
		if err != nil {

		}
	}
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/notes" && !strings.HasPrefix(r.URL.Path, "/notes/") {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodPost:
		s.NotesPost(w, r)
	case http.MethodGet:
		s.NotesGet(w, r)
	case http.MethodDelete:
		s.NotesDelete(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
