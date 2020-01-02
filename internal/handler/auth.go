package handler

import (
	"context"
	"encoding/json"
	"github.com/vinchauhan/two-f-gates/internal/service"
	"log"
	"net/http"
	"strings"
)

type loginInput struct {
	Email string
	Username string
}
func (h *handler) withAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimSpace(r.URL.Query().Get("auth_token"))

		if token == "" {
			if a := r.Header.Get("Authorization"); strings.HasPrefix(a, "Bearer ") {
				token = a[7:]
			}
		}
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}
		log.Printf("Token is : %s", token)
		uid, err := h.AuthUserIDFromToken(token)

		log.Printf("UID is : %s", uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, service.KeyAuthUserID, uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func (h *handler) devLogin(w http.ResponseWriter, r *http.Request) {
	var in loginInput
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := h.DevLogin(r.Context(), in.Email)
	//if err == service.ErrUnimplemented {
	//	http.Error(w, err.Error(), http.StatusNotImplemented)
	//	return
	//}

	if err == service.ErrInvalidEmail {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err == service.ErrUserNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	respond(w, out, http.StatusOK)
}
