package handler

import (
	"mime"
	"net/http"
	"net/url"

	"github.com/matryer/way"
	"github.com/vinchauhan/two-f-gates/internal/service"
)

type handler struct {
	*service.Service
}


//New Returns a new mux router
func New(s *service.Service, origin url.URL) http.Handler {
	h := handler{s}
	api := way.NewRouter()
	api.HandleFunc("POST", "/user", h.createUser)
	api.HandleFunc("POST", "/twillio", h.twillioHandler)
	api.HandleFunc("POST", "/passcodes", h.generatePasswords)
	api.HandleFunc("GET", "/passcodes", h.getPasscodes)
	api.HandleFunc("POST", "/dev_login", h.devLogin)
	// api.HandleFunc("POST", "/passcode/validate", h.validatePasscode)

	mime.AddExtensionType(".js", "application/javascript; charset=utf-8")

	//Custom fileServer with own fileSystem
	fs := http.FileServer(&spaFileSystem{http.Dir("web/static/")})

	if origin.Hostname() == "localhost" {
		fs = withoutCache(fs)
	}

	r := way.NewRouter()
	r.Handle("*", "/api...", http.StripPrefix("/api", h.withAuth(api)))
	r.Handle("GET", "/...", fs)

	return r

}
