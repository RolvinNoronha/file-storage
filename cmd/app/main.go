package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RolvinNoronha/fileupload-backend/internal/db"
	"github.com/RolvinNoronha/fileupload-backend/internal/env"
)

func main() {
	env.InitializeEnv();
	db, err := db.InitializeDb();

	if (err != nil) {
		log.Fatal(err);
	}

	port := os.Getenv("PORT")
	cfg := config{
		addr: fmt.Sprintf(":%s", port),
		db: db,
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
