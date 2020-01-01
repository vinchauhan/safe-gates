package handler

import (
	"github.com/vinchauhan/two-f-gates/internal/service"
	"net/http"
)

import "encoding/json"

type createUserInput struct {
	Email    string
	Username string
}

func (h *handler) createUser(w http.ResponseWriter, r *http.Request) {
	var in createUserInput
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.CreateUser(r.Context(), in.Email, in.Username)

	if err == service.ErrUserNameExists {
		respond(w, service.ErrUserNameExists.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


