run:
	go run ./cmd/fileupload/main.go

goose_up:
	cd ./migrations && goose postgres postgresql://admin:password@localhost:5432/filestorage up

goose_down:
	cd ./migrations && goose postgres postgresql://admin:password@localhost:5432/filestorage down

