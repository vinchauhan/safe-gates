package main

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/vinchauhan/two-f-gates/internal/handler"
	"github.com/vinchauhan/two-f-gates/internal/service"
	"log"
	"net/http"
)

const (
	port = 3000
)

func main() {

	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	s := service.New(db)
	h := handler.New(s)
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Accepting connection on port %d", port)
	log.Fatalf("Error starting the sever %s", http.ListenAndServe(addr, h))
}
