package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vinchauhan/two-f-gates/internal/handler"
)

const (
	port = 3000
)

func main() {

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Accepting connection on port %d", port)
	log.Fatalf("Error starting the sever %s", http.ListenAndServe(addr, handler.New()))
}
