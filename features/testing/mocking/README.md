## Mocking

Mocking is a fundamental aspect of unit testing in software development. In Golang, mocking helps isolate the code under test by replacing dependencies with mock implementations.


### What is Mockgen?
mockgen is a command-line tool provided by the GoMock library that generates mock implementations of interfaces. It is widely used in the Golang community to create mocks for unit testing. The tool automates the process of writing mock objects, which saves time and reduces boilerplate code.

### How to generate mocking

```
mockgen -destination=./features/testing/mocking/mock/clientMock.go -package=mocks samanthajayasinghe.github.io/golang/features/testing/mocking PagerDutyClient
```

### How to run test 
```
go test -v samanthajayasinghe.github.io/golang/features/testing/mocking
```
