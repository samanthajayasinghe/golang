# Go Version : 1.21
release date : 2023-08-08
release notes : https://go.dev/doc/go1.21

## New Features
### Changes to the language
1. The new functions min and max compute the smallest (or largest, for max) value of a fixed number of given arguments.
2. The new function clear deletes all elements from a map or zeroes all elements of a slice.

### Tools

Go 1.21 adds improved support for backwards compatibility and forwards compatibility in the [Go toolchain](https://go.dev/doc/toolchain).

How to run in older version 

```
GOTOOLCHAIN=go1.20 go run versions/1.21/features/min/main.go
```
