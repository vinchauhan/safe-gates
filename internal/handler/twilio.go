package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/matryer/way"
)

//<Record timeout="5"></Record>
const (
	welcomeMessage = `<Response>
						  <Say voice="alice">Welcome to Chauhan Family. Please enter passcode</Say>
						  <Gather finishOnKey="#" timeout="5"></Gather>
					  </Response>`

	authInprogressMsg = `<Response>
		<Play>Your password entered is correct : %s</Play>
	</Response>`
)

//New Returns a new mux router
func New() http.Handler {
	api := way.NewRouter()
	api.HandleFunc("POST", "/", twillioHandler)
	return api
}

func twillioHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/xml")

	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	for k, v := range r.PostForm {
		fmt.Printf("%s : %s\n", k, v)
	}

	passcode := r.FormValue("Digits")

	if passcode == "" {
		fmt.Fprintf(w, welcomeMessage)
		return
	}
	fmt.Fprintf(w, authInprogressMsg, passcode)
	b, _ := httputil.DumpRequest(r, true)
	log.Printf("request is %s", b)

	fmt.Fprintf(w, welcomeMessage)
}
