package server

import (
	"net/http"
)

func (s *Server) routes() {
	http.HandleFunc("/notes", s.ServeHTTP)
	http.HandleFunc("/notes/", s.ServeHTTP)
}
