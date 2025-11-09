package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RolvinNoronha/fileupload-backend/internal/aws"
	"github.com/RolvinNoronha/fileupload-backend/internal/db"
	"github.com/RolvinNoronha/fileupload-backend/internal/env"
)

func main() {
	env.InitializeEnv()
	ps := db.InitializePostgres()
	es := db.InitializeElasticSearch()
	client := aws.InitializeAws()

	port := os.Getenv("PORT")
	cfg := config{
		addr:   fmt.Sprintf(":%s", port),
		ps:     ps,
		es:     es,
		client: client,
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
