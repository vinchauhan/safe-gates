package handler

import (
	"fmt"
	"net/http"
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

func (h *Handler) twillioHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/xml")

	for k, v := range r.PostForm {
		fmt.Printf("%s, %s", k, v)
	}

	rec := r.FormValue("RecordingUrl")

	if rec == "" {
		fmt.Fprintf(w, welcomeMessage)
		return
	}
	fmt.Fprintf(w, authInprogressMsg, rec)
}
