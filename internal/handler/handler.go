package handler

import (
	"github.com/matryer/way"
	"github.com/vinchauhan/two-f-gates/internal/service"
	"net/http"
)

type handler struct {
	*service.Service
}
//New Returns a new mux router
func New(s *service.Service) http.Handler {
	h := handler{s}
	api := way.NewRouter()
	api.HandleFunc("POST", "/", h.twillioHandler)
	api.HandleFunc("POST", "/passcodes", h.generatePasswords)
	api.HandleFunc("GET", "/passcodes", h.getPasscodes)
	return api
}
