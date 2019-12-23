package handler

import (
	"net/http"
)

func (h *handler) generatePasswords(resp http.ResponseWriter, req *http.Request) {
	//TODO: In the future find more secure way to get passcodes
	//Sets the passcodes received as slice of string in the db
	defer req.Body.Close()

	codes, err := h.GeneratePasscodes()
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	respond(resp, codes, http.StatusOK)
}

func (h *handler) getPasscodes(resp http.ResponseWriter, req *http.Request) {
	//TODO: In the future find more secure way to get passcodes
	//Sets the passcodes received as slice of string in the db
	out, err := h.GetPassCodes()
	if err != nil {
		http.Error(resp, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	respond(resp, out, http.StatusOK)
}
