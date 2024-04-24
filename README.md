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
cd data
go test -v . --tags integration --count=1
go test -cover . -v --tags integration --count=1
go test -cover . --tags integration --count=1


go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

```
cd celeritas
go get github.com/joho/godotenv

go get -u github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware

go get github.com/CloudyKit/jet/v6

go get github.com/alexedwards/scs/v2

go get github.com/jackc/pgconn
go get github.com/jackc/pgx/v4
go get github.com/jackc/pgx/stdlib

cd myapp
go get github.com/upper/db/v4/adapter/postgresql
go get github.com/upper/db/v4/adapter/mysql
go mod vendor 

#test Models
go get github.com/DATA-DOG/go-sqlmock

go get github.com/ory/dockertest/v3

go get github.com/fatih/color

cd celeritas
go get github.com/golang-migrate/migrate/v4

go get github.com/golang-migrate/migrate/v4/database/mysql

go get github.com/golang-migrate/migrate/v4/database/postgres

go get github.com/golang-migrate/migrate/v4/database/file

```
```
https://upper.io/v4/
```


### run below code with root user to see percent of coverage 
```
go tool cover -html=coverage.out -o coverage.html && google-chrome --no-sandbox coverage.html
```

```
#!/usr/bin/zsh

# Allow user "chrome" to access the X server
xhost +si:localuser:chrome

# Run Google Chrome as user "chrome" with the specified display
sudo -H -u chrome DISPLAY=:0 /usr/bin/google-chrome


# useradd chrome -s /usr/bin/zsh -d /home/chrome -m 
#!/usr/bin/zsh
xhost +
sudo -H -u chrome DISPLAY=:0 google-chrome

#chmod a+x /root/Doccuments/chrome.sh

```
