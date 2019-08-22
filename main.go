package main

import (
	"log"
	"net/http"

	"github.com/matryer/way"
	"github.com/vinchauhan/two-f-gates/internal/handler"
)

func main() {

	router := way.NewRouter()
	h := handler.New()

	router.HandleFunc("POST", "/", handler.MainHandler)
	log.Fatalln(http.ListenAndServe(":8080", router))
}
