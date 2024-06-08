#run: build
	#@./bin/api

#build:
	#@go build -o bin/api

#test:
	#@go test -v ./...


# Run command
run: build
	@DB_USER=root DB_PASSWORD=m4c3nz29 DB_HOST=localhost DB_NAME=projectmanager ./bin/api

# Build command
build:
	@go build -o bin/api

# Test command
test:
	@go test -v ./...
