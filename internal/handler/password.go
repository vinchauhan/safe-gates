package handler

import (
	"github.com/matryer/way"
	"github.com/vinchauhan/two-f-gates/internal/service"
	"net/http"
)

func (h *handler) generatePasswords(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	username := way.Param(ctx, "username")
	//Sets the passcodes received as slice of string in the db
	defer req.Body.Close()

	codes, err := h.GeneratePasscodes(ctx, username)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	respond(resp, codes, http.StatusOK)
}

func (h *handler) getPasscodes(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	username := way.Param(ctx, "username")

	//Sets the passcodes received as slice of string in the db
	out, err := h.GetPassCodes(ctx, username)
	if err == service.ErrUnauthenticated {
		http.Error(resp, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	if err != nil {
		http.Error(resp, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	respond(resp, out, http.StatusOK)
}
