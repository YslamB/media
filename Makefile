include .env

dev:
	@go run main.go
db:
	@echo "Initializing ussat database..."
	@psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d postgres \
		-f ./database/init/init.sql
	@psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d ussat \
		-f ./database/init/create.sql
	@psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d ussat \
		-f ./database/init/insert.sql
	@psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d ussat \
		-f ./database/init/functions.sql
	@echo "Has been successfully created"
build:
	@echo "Building the app, please wait..."
	@go build -o ./bin/ussat main.go
	@echo "Done."
build-cross:
	@echo "Bulding for windows, linux and macos (darwin m2), please wait..."
	@GOOS=linux GOARCH=amd64 go build -o ./bin/ussat-linux main.go
	@GOOS=darwin GOARCH=arm64 go build -o ./bin/ussat-macos main.go
	@GOOS=windows GOARCH=amd64 go build -o ./bin/ussat-windows main.go
	@echo "Done."
deploy:
	scp ./bin/ussat-linux ussa@172.16.19.104:/var/www/ussa/
deploy-db:
	scp -r ./database/init ussa@172.16.19.104:/var/www/ussa/
