package main

import (
	"fmt"
	"log"
	"os"

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

	port := os.Getenv("PORT")
	cfg := config{
		addr: fmt.Sprintf(":%s", port),
		db: db,
		jwtSecret: []byte(os.Getenv("JWT_SECRET")),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
