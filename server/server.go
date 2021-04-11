package server

import (
	"net/http"
	"user/handler"
	"user/middleware"
)

type Server struct {
	Router *http.ServeMux
}

func (s *Server) InitRoute(h *handler.Handler) {
	s.Router.HandleFunc("/user/profile", middleware.JSONandCORS(h.UserProfile))
	s.Router.HandleFunc("/microservice/name", middleware.JSONandCORS(h.MicroserviceName))
}
