package main

import (
	"log"
	"net/http"
	"time"

	"github.com/RolvinNoronha/fileupload-backend/internal/routes"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/elastic/go-elasticsearch/v9"
	"gorm.io/gorm"
)

type config struct {
	addr   string
	ps     *gorm.DB
	es     *elasticsearch.Client
	client *s3.Client
}

type application struct {
	config config
}

func (app *application) mount() http.Handler {

	router := routes.NewRouter(app.config.ps, app.config.client)
	return router

	/*
		mux := http.NewServeMux()
		userRepo := user.NewRepository(app.config.db);
		userService := user.NewService(userRepo);
		userHandler := user.NewHandler(userService);

		mux.HandleFunc("POST /register", userHandler.CreateUser);
		mux.HandleFunc("/protected", userHandler.Protected);
		mux.HandleFunc("/refresh", userHandler.Refresh);

		return mux
	*/
}

func (app *application) run(router http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server listening on port %s", app.config.addr)

	return srv.ListenAndServe()
}
