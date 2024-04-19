package main

import (
	"fmt"
	"github.com/sirupsen/logrus"

	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	env  string
}
type application struct {
	config config
}

var log = logrus.New()

func main() {
	var cfg config

	log.SetFormatter(&logrus.JSONFormatter{})
	logFile, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error("Failed to open log file: ", err)
		return
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	app := &application{
		config: cfg,
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 30,
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
