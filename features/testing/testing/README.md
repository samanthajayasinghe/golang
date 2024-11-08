# Go Testing
Go's built-in support for unit testing makes it easier to test as you go. Specifically, using naming conventions, Go's testing package, and the go test command, you can quickly write and execute tests.

## How to execute test
```
go test -v
go test -v -parallel
```

## test coverage 
```
go test -coverprofile=cover.out
go tool cover -html=cover.out

```
