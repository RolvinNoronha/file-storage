package main

import (
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type application struct {
	config config
}

type config struct {
	addr string
	db *gorm.DB
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ser", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	return mux
}

func (app *application) run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server listening on port %s", app.config.addr)

	return srv.ListenAndServe()
}
