# Basic API with Go Gin, Jwt and Gorm

Basic CRUD API implemented with go Go [gin](github.com/gin-gonic/gin) and [gorm](https://gorm.io/). The API allows you to perform basic CRUD operations on a database, as well as create and login users using `jwt`.

Current implementation is using CockroachDB with a postgres connection.

## Tools and Libraries

```
go get github.com/githubnemo/CompileDaemon also go install github.com/githubnemo/CompileDaemon
go get github.com/joho/godotenv
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

/Users/barnatest/go/bin/CompileDaemon -command="./go_crud" `run the compile daemon`
```

## Installation

1. Clone the repository
2. Navigate to the project directory
3. Install the dependencies: `go mod download`

## Usage

1. Start the server: `go run main.go`
2. Set server -> : `http://localhost:3000`
3. Use an API testing tool like Postman to send HTTP requests to the API endpoints.
4. Check out this cool tool for database visualization [here](https://www.beekeeperstudio.io/get-community)

<!-- https://www.beekeeperstudio.io/get-community -->
