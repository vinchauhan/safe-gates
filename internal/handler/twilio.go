package handler

import (
	"fmt"
	"log"
	"net/http"
)

//<Record timeout="5"></Record>
const (
	welcomeMessage = `<?xml version="1.0" encoding="UTF-8"?>
						<Response>
						  <Say voice="alice">Please enter passcode</Say>
						  <Gather finishOnKey="#" timeout="5"></Gather>
					  </Response>`

	accessGrantedMsg = `<?xml version="1.0" encoding="UTF-8"?>
							<Response>
								<Say voice="alice">Access Granted</Say>
								<Play digits="9"></Play>
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
		res, err := h.validatePasscode(passcode)
		if err != nil {
			fmt.Fprintf(w, applicationError, err)
			return
		}
		if res {
			fmt.Fprintf(w, accessGrantedMsg)
		} else {
			fmt.Fprintf(w, accessDenied)
		}
	}
}

func (h *handler) validatePasscode(passcode string) (bool, error) {

	//ValidatePasscode will check if password exists in DB
	b, err := h.ValidatePasscode(passcode)
	if err != nil {
		log.Fatal(err)
	}
	return b, nil
}
