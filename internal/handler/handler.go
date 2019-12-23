package handler

import (
	"net/http"

	"github.com/matryer/way"
	"github.com/vinchauhan/two-f-gates/internal/service"
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
	// api.HandleFunc("POST", "/passcode/validate", h.validatePasscode)
	return api
}
