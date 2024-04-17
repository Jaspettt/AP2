package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"task1/internal/mysql"
	"time"
)

type config struct {
	port int
	env  string
}
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "enviroment", "development", "API DEVELOPMENT")
	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	dsn := "jaspet:1337@tcp(localhost:3306)/VinylsGolang"
	db, err := mysql.Connect(dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 30,
	}
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)

}
