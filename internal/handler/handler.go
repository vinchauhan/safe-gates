package handler

import (
	"github.com/matryer/way"
	"github.com/vinchauhan/two-f-gates/service"
	"net/http"
)

type Handler struct {
	*service.Service
}

func New(s *service.Service) http.Handler {

	h := &Handler{s}

	api := way.NewRouter()
	api.HandleFunc("*", "/", h.twillioHandler)
	//api.HandleFunc("POST", "/login", h.loginUser)
	//api.HandleFunc("POST", "/users", h.createUser)
	//api.HandleFunc("POST", "/users/:username/toggle_follow", h.toggleFollow)
	//api.HandleFunc("GET", "/auth_user", h.authUser)
	//api.HandleFunc("GET", "/users/:username", h.getUser)

	//r := way.NewRouter()
	//r.Handle("*", "/api...", http.StripPrefix("/api", h.withAuthMiddleware(api)))
	return api
}
