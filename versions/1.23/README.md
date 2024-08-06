# Go Version : 1.23
expected release date : 2024-08
release notes : https://tip.golang.org/doc/go1.23

## New Features
### Changes to the language
1. The “range” clause in a “for-range” loop now accepts iterator functions of the following types
```
func(func() bool)
func(func(K) bool)
func(func(K, V) bool)
```

### Tools
1. `go telemetry` [Go telemetry](https://tip.golang.org/doc/telemetry) is a way for Go toolchain programs to collect data about their performance and usage. 
