package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vinchauhan/two-f-gates/internal/handler"
	"github.com/vinchauhan/two-f-gates/internal/service"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"
)

const (
	port = 3000
)

func main()  {
	godotenv.Load()
	if err := run(); err != nil {
		log.Printf("Error is %v\n", err)
		log.Fatalln(err)
	}
}

func run() error {
	var (
		port, _  = strconv.Atoi(env("PORT", "3000"))
		originStr = env("ORIGIN", fmt.Sprintf("http://localhost:%d", port))
		dbURL = env("RETHINK_DB_URL", "localhost:28015")
		tokenKey     = env("TOKEN_KEY", "supersecretkeyyoushouldnotcommit")
	)

	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Println("\nDont forget to set tokenKey in environment or while starting the app")
	}

	//override the setting in environment to start server
	flag.IntVar(&port, "port", port, "Port which the server will run on")
	flag.StringVar(&originStr, "origin", originStr, "URL origin of the service")
	flag.StringVar(&tokenKey, "secret", tokenKey, "Token key or secret")
	flag.StringVar(&dbURL, "db", dbURL, "Database URL")
	flag.Parse()

	origin, err := url.Parse(originStr)
	if err != nil || !origin.IsAbs() {
		return errors.New("invalid url origin")
	}

	log.Printf("BEGIN")
	session, err := r.Connect(r.ConnectOpts{
		//Address: os.Getenv("RETHINK_DB_URL"),
		Address: dbURL,
	})

	if err != nil {
		log.Fatal(err)
	}
	//Recreate tables
	err = r.DB("test").TableDrop("passcodes").Exec(session)
	err = r.DB("test").TableCreate("passcodes").Exec(session)

	err = r.DB("test").TableDrop("users").Exec(session)
	err = r.DB("test").TableCreate("users").Exec(session)

	s := service.New(session, *origin, tokenKey)
	addr := fmt.Sprintf(":%d", port)
	server := http.Server{
		Addr:              addr,
		Handler:           handler.New(s, *origin),
		ReadHeaderTimeout: time.Second * 5,
		ReadTimeout:       time.Second * 15,
		WriteTimeout:      time.Second * 15,
		IdleTimeout:       time.Second * 30,
	}

	errs := make(chan error, 2)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, os.Kill)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			errs <- fmt.Errorf("could not shutdown server: %v", err)
			return
		}

		errs <- ctx.Err()
	}()
	if err = server.ListenAndServe(); err != nil {
		log.Fatalln("Could not start server")
	}
	//go func() {
	//	log.Printf("Accepting connection on port %d", port)
	//	log.Fatalf("Starting server at origin %s\n", origin)
	//	if err = server.ListenAndServe(); err != http.ErrServerClosed {
	//		log.Printf("Error is %v\n", err)
	//		errs <- fmt.Errorf("could not listen and serve: %v", err)
	//		return
	//	}
	//
	//	errs <- nil
	//}()
	return <-errs
}

func env(key , defaultValue string) string {
	s, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return s
	
}
