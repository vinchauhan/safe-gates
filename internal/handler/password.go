package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

//Passcodes struct stores passcodes
type Passcodes struct {
	passcodes []string
}

func (h *handler) generatePasswords(resp http.ResponseWriter, req *http.Request)  {
	//TODO: In the future find more secure way to get passcodes
	//Sets the passcodes received as slice of string in the db
	var in Passcodes
	defer req.Body.Close()

	//Decode the passcodes Struct
	if err := json.NewDecoder(req.Body).Decode(&in); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	codes , err := h.GeneratePasscodes(in.passcodes)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	respond(resp, codes, http.StatusOK)
}

func (h *handler) getPasscodes(resp http.ResponseWriter, req *http.Request) {
	//TODO: In the future find more secure way to get passcodes
	//Sets the passcodes received as slice of string in the db
	var out []string
	defer req.Body.Close()
	out, err := h.GetPassCodes()
	if err != nil {
		panic(err)
	}
	log.Printf("Passcodes are : %s", out)
}
