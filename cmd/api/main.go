package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "0.0.1"

type config struct {
	port int
	env  string
	db struct {dsn string}
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
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN") , "PostgreSQL DSN")
	flag.Parse()

	//Logger to write to standard output
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	logger.Printf("dsn: %v",cfg.db.dsn)

	db, err := openDB(cfg)
    if err != nil {
        logger.Fatal(err)
    }

	defer db.Close()
	logger.Printf("database connection pool established")

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
	err = srv.ListenAndServe()
	logger.Fatal(err)

}


func openDB(cfg config) (*sql.DB, error) {
    // Use sql.Open() to create an empty connection pool, using the DSN from the config
    // struct.
    db, err := sql.Open("postgres", cfg.db.dsn)
    if err != nil {
        return nil, err
    }

    // Create a context with a 5-second timeout deadline.
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Use PingContext() to establish a new connection to the database, passing in the
    // context we created above as a parameter. If the connection couldn't be
    // established successfully within the 5 second deadline, then this will return an
    // error.
    err = db.PingContext(ctx)
    if err != nil {
        return nil, err
    }

    // Return the sql.DB connection pool.
    return db, nil
}
