package main

import (
	// "fmt"
	"log"

	"github.com/RolvinNoronha/fileupload-backend/internal/db"
	"github.com/RolvinNoronha/fileupload-backend/internal/env"
	// "net/http"
)

func main() {
	env.InitializeEnv();
	db, err := db.InitializeDb();

	if (err != nil) {
		log.Fatal(err);
	}

	cfg := config{
		addr: ":8080",
		db: db,
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
