package server

import (
		"github.com/matryer/way"
		"net/http"
	"fmt"
	)

type Server struct {
	db     string
	router *way.Router
	email  string

}

func (s *Server) routes() {
	s.router.HandleFunc("GET", "/api/", s.handleAPI())
	//s.router.HandleFunc("GET", "/login/:username/:password", s.handleLogin())
	//s.router.HandleFunc("GET", "/", s.handleIndex())
}


func (s *Server) handleAPI() http.HandlerFunc  {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Wellcome to my API web page!!")
	}
}

func NewServer() *Server{
	return  &Server{
		db :"",
		router : way.NewRouter(),
	}
}