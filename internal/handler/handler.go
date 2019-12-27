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

	//Custom fileServer with own fileSystem
	fs := http.FileServer(&spaFileSystem{http.Dir("web/static")})

	
	r := way.NewRouter()
	r.Handle("*", "/api...", http.StripPrefix("/api", h.withAuth(api)))
	r.Handle("GET", "/...", fs)
	
	return r

}
