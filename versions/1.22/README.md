# Go Version : 1.22
release date : 2024-02-06
release notes : https://tip.golang.org/doc/go1.22

## New Features
### Changes to the language
1. “For” loops may now range over integers.
2. Previously, the variables declared by a “for” loop were created once and updated by each iteration. In Go 1.22, each iteration of the loop creates new variables

### Tools
1. The directory is created by `go work vendor`, and used by build commands when the -mod flag is set to vendor, which is the default when a workspace vendor directory is present.

2. `go get` is no longer supported outside of a module 
