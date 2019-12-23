package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

//<Record timeout="5"></Record>
const (
	welcomeMessage = `<Response>
						  <Say voice="alice">Welcome to Chauhan Family. Please enter passcode</Say>
						  <Gather finishOnKey="#" timeout="5"></Gather>
					  </Response>`

	accessGrantedMsg = `<Response>
						 	<Say voice="alice">Access Granted</Say>
						 </Response>`

	accessDenied = `<Response>
						 	<Say voice="alice">Access Denied.!</Say>
						 </Response>`

	applicationError = `<Response>
						 	<Say voice="alice">Application Error occurred : %v</Say>
						 </Response>`
)

func (h *handler) twillioHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/xml")

	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	for k, v := range r.PostForm {
		fmt.Printf("%s : %s\n", k, v)
	}
	passcode := r.FormValue("Digits")
	log.Printf("Entered Passcode is : %s", passcode)
	if passcode == "" {
		fmt.Fprintf(w, welcomeMessage)
	} else {
		//Which means the request from Twillio has Digits pressed by the human.
		// Validate the passcode digits.
		res, err := validatePasscode(passcode)
		if err != nil {
			fmt.Fprintf(w, applicationError, err)
			return
		}
		if res {
			fmt.Fprintf(w, accessGrantedMsg)
		}
		//b, _ := httputil.DumpRequest(r, true)
		//log.Printf("request is %s", b)
		return
	}
	//fmt.Fprintf(w, welcomeMessage)
}

func validatePasscode(passcode string) (bool, error) {
	return true, errors.New("Password cannot be validated")
}
