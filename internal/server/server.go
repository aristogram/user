package server

import (
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Start() {
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "ты лох", 200)
	})

	err := http.ListenAndServe("localhost:14888", s.mux)
	if err != nil {
		log.Fatalln("server error:", err)
	}
}
