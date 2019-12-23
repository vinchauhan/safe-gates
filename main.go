package main

import (
	"fmt"
	"github.com/vinchauhan/two-f-gates/internal/handler"
	"github.com/vinchauhan/two-f-gates/internal/service"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"log"
	"net/http"
	"os"
)

const (
	port = 3000
)

func main() {

	log.Printf("BEGIN")
	session, err := r.Connect(r.ConnectOpts{
		Address: os.Getenv("RETHINK_DB_URL"),
	})

	if err != nil {
		log.Fatal(err)
	}
	err = r.DB("test").TableDrop("passcodes").Exec(session)
	err = r.DB("test").TableCreate("passcodes").Exec(session)

	s := service.New(session)
	h := handler.New(s)
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Accepting connection on port %d", port)
	log.Fatalf("Error starting the sever %s", http.ListenAndServe(addr, h))
}
