package handler

import (
	"github.com/matryer/way"
	"github.com/vinchauhan/two-f-gates/service"
	"net/http"
)

const (
	welcomeMessage = `<Response>
	<Say voice="alice">Please say your password after the beep.</Say>
	<Record time="7"></Record>
	</Response>`

	authInprogressMsg = `<Response>
	<Say voice="alice">Validating Password</Say>
	</Response>`
)

type handler struct {
	*service.Service
}
func New(s *service.Service) http.Handler{
	h := &handler{s}

	api := way.NewRouter()

	api.HandleFunc("POST", "/login", h.loginUser)
	api.HandleFunc("POST", "/users", h.createUser)
	api.HandleFunc("POST", "/users/:username/toggle_follow", h.toggleFollow)
	api.HandleFunc("GET", "/auth_user", h.authUser)
	api.HandleFunc("GET", "/users/:username", h.getUser)

	r := way.NewRouter()
	r.Handle("*", "/api...", http.StripPrefix("/api", h.withAuthMiddleware(api)))
	return r
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/xml")

	w.Write([]byte(``))
}
