## Using Go Modules
## https://blog.golang.org/using-go-modules

Go's new dependency management system makes dependency version information:
- explicit
- easier to manage

Module: a collection of Go packages stored in a file tree with a go.mod file as its root.

go.mod file defines:
- module path, the import path used for the root directory
- dependency rquirements, other modules needed for a successful build
  - each dependency requirement is written as a module path and a specific semantic version



## Follow these steps for better understanding
- Creating a new module
- Adding a dependency
- Upgrading dependencies
- Adding a dependency on a new major version
- Upgrading a dependency on a new major version
- Removing unused dependencies

### Step 1: Creating a new module

> mkdir /home/gregory/gopher
> cd /home/gregory/gopher
> tree
├── hello.go
└── hello_test.go

> go test
PASS
ok      _/home/gregory/gopher   0.001s
#### because we are outside $GOPATH and outside any module, 
#### go command creates a fake import path based on directory name

> go mod init example.com/hello
> tree
├── go.mod
├── hello.go
└── hello_test.go

> go test
PASS
ok      example.com/hello       0.001s
#### go command knows the go module

- go.mod file only appears in the root of the module

### Step 2: Adding a dependency

- The primary motivation for `go module` was to improve the experience of using code written by other developers

#### modify hello.go to import "rsc.io/quote"
#### run the test again

> go test
PASS
ok      example.com/hello       0.002s
> cat go.mod 
module example.com/hello

go 1.14

require rsc.io/quote v1.5.2

> go list -m all
example.com/hello
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0

#### note above, `main module` is always first in the list
#### dependencies, and their depdencies, follow

#### `go` command also maintains a file named `go.sum` containing the expected cryptographic hashes of the content of specific module versions
> cat go.sum

#### both `go.mod` and `go.sum` should be checked into version control




### Step 3: Upgrading a dependency

> go list -m all
example.com/hello
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0

> go get golang.org/x/text
> go test
> go list -m all
example.com/hello
golang.org/x/text v0.3.3
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
> cat go.mod
module example.com/hello

go 1.14

require (
        golang.org/x/text v0.3.3 // indirect
        rsc.io/quote v1.5.2
)

> go get rsc.io/sampler

#### this version is incompatible, specify a specific version:

> go get rsc.io/sampler@v1.3.1






https://blog.golang.org/using-go-modules



### Step 4: Adding a dependency on a new major version
- @TODO
### Step 5: Upgrading a dependency on a new major version
- @TODO
### Step 6: Removing unused dependencies
- @TODO





### Conclusion
- Go modules are the future of dependency management in Go.
- Module functionality is now available in all supported Go versions.

This post introduced these workflows using Go modules:
- go mod init creates a new module, initializing the go.mod file that describes it
- go build, go test, and other package-building commands add new dependencies to go.mod as needed
- go list -m all prints the current module's dependencies
- go get changes the required version of a dependency (or adds a new dependency)
- go mod tidy removes unused dependencies

We encourage you to start using modules in your local development and to add go.mod and go.sum files to your projects.