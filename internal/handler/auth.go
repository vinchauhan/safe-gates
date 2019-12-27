package handler

import (
	"net/http"
)

func (h *handler) withAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		
	})
}