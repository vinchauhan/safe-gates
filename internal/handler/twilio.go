package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/matryer/way"
)

const (
	welcomeMessage = `<Response>
		<Say voice="alice">Please say your password after the beep.</Say>
		<Record timeout="5"></Record>
	</Response>`

	authInprogressMsg = `<Response>
		<Play>%s</Play>
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

	// for k, v := range r.PostForm {
	// 	fmt.Printf("%s, %s", k, v)
	// }

	// rec := r.FormValue("RecordingUrl")

	// if rec == "" {
	// 	fmt.Fprintf(w, welcomeMessage)
	// 	return
	// }
	// fmt.Fprintf(w, authInprogressMsg, rec)
	b, _ := httputil.DumpRequest(r, true)
	log.Printf("request is %s", b)

	fmt.Fprintf(w, welcomeMessage)
}
