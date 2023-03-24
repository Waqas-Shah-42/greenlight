package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "0.0.1"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	fmt.Println("Starting server")

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production)")
	flag.Parse()

	//Logger to write to standard output
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Creating application and passing logger and configuration to it.
	app := &application{
		config: cfg,
		logger: logger,
	}

	// intiating server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),

		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
