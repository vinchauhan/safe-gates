package main

import (
	"fmt"
	"github.com/vinchauhan/two-f-gates/service"
	"log"
	"net/http"

	"github.com/vinchauhan/two-f-gates/internal/handler"
)

const (
	databaseURL = "postgresql://root@127.0.0.1:26257/two-f-gates?sslmode=disable"
	port        = 3000
)

func main() {

	s := service.New("hello")

	h := handler.New(s)

	addr := fmt.Sprintf(":%d", port)

	log.Printf("Accepting connection on port %d", port)

	//if err = http.ListenAndServe(addr, h); err != nil {
	//	log.Fatalf("Failed to start sever: %v\n", err)
	//}
	log.Fatalf("Error starting the sever %s", http.ListenAndServe(addr, h))
}
