# go-laravel
go-laravel

Initial Project 
```
├── go-laravel
│   ├── celeritas
│   └── myapp

```
```
cd celeritas
go mod init github.com/Moji00f/celeritas
cd ../myapp
go mod init myapp
```

in `myapp` and in `go.mod` add below config 
```
replace github.com/Moji00f/celeritas => ../celeritas
```

when do change in celeritas you should update myapp like below
```
go get github.com/Moji00f/celeritas 
go: added github.com/Moji00f/celeritas v0.0.0-00010101000000-000000000000

go get -u github.com/Moji00f/celeritas 
```

insted fetching data
```
go mod vendor
```

Create `Makefile` in root level of `myapp`
```
BINARY_NAME=celeritasApp

build:
	@go mod vendor
	@echo "Building Celeritas..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Celeritas built!"

run: build
	@echo "Starting Celeritas..."
	@./tmp/${BINARY_NAME} &
	@echo "Celeritas started!"

start: run

stop:
	@echo "Stopping Celeritas..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Celeritas!"

restart: stop start

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "CLeaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"
```
[install GoDotEnv](https://github.com/joho/godotenv)

```
cd celeritas
go get github.com/joho/godotenv
go get -u github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
```
