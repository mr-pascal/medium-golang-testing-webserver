# medium-golang-testing-webserver

Example project for testing a web server ([Fiber](https://github.com/gofiber/fiber)) in Go.

### Running tests

```sh
## Go to service directory
cd src/service

## Run tests
go test

## Run tests with verbose output
go test -v
```


### Test Coverage

```sh

## Get test coverage
go test -cover

## Generate test coverage HTML file
go test -cover -coverprofile=./tmp/c.out
go tool cover -html=./tmp/c.out -o ./tmp/coverage.html 
```