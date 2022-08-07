# Quiz-Master
[![wakatime](https://wakatime.com/badge/user/90095cec-bff0-4e71-9e19-19fbac89fb11/project/9a69d0dd-130c-49b6-93ab-2a629bd2b36b.svg)](https://wakatime.com/badge/user/90095cec-bff0-4e71-9e19-19fbac89fb11/project/9a69d0dd-130c-49b6-93ab-2a629bd2b36b)
![Golang CI Lint](https://github.com/jidniilman/Quiz-Master/actions/workflows/golangci-lint.yml/badge.svg)

----
Quiz Master is Quipper’s Technical Exam for Backend Engineer - Web

## Overview
- This project is tracked by WakaTime
- This project built with Go 1.19
- This project init by Linux (Ubuntu WSL2 on Windows)
- This project developed and coded with Windows 10
- Detailed time elapsed to write this project can be found [here (WakaTime)](https://wakatime.com/@deraven/projects/qstzzolwiz?start=2022-08-02&end=2022-08-08) 

## Todo
- [x] App Functionality
- [x] Push to Github for first time
- [ ] Cleanup Readme
- [x] Write Unit Tests
- [ ] Write UNIX Executable Script `/bin/setup` and `/bin/quiz_master`
- [ ] Implement GitHub Actions
- [x] Write proper documentation
- [ ] Implement Golang CI Lint
- [x] Write Benchmark

## Tech used
- Go Modules
- Makefile (todo)

## Library Used
Use own logic as much as possible. everything else is self-made.
- [Testify](https://github.com/stretchr/testify) for Unit Testing purpose

## Coding Convention
Golang is strict by default. But we are using some guidelines from :
- [Effective Go](https://go.dev/doc/effective_go)
- [CockroachDB Engineering Standards](https://wiki.crdb.io/wiki/spaces/CRDB/pages/181371303/Go+Golang+coding+guidelines)
- [Golang CI Lint](https://golangci-lint.run/)
- GoLand IDE formatting

## Unit Tests
Mainly for our unit testing, we are using table driven test approach. But we are using little sample to do our tests.

Run all unit tests with `go test ./...` or `go test -v ./...` with verbose output.
You can run individual unit test with `go test -v <package-name>`.

Package Name (main entrypoint are excluded):
```
jidniilman/quiz-master/internal/cli      # Handling CLI application and execution
jidniilman/quiz-master/internal/question # Our question model
jidniilman/quiz-master/pkg/command       # Our data communication protocol
jidniilman/quiz-master/pkg/utils         # Some useful helpers
```
To see full package list, use `go list ./...`

## Benchmark
Only `/pkg/command` and `/pkg/utils` that have benchmark.

Run benchmark with:
```
go test -v ./... -bench=. -benchmem
```

## Project Layout/Structure
In this project, I'm using the layout standards from [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
I remove some of unused folders such as: `/api`, `/web`, `/deployments`, `/test` 

### `/cmd`
Main applications for this project. The directory name for each application match the name of the executable (e.g., `/cmd/quiz_master`).
It's common to have a small main function that imports and invokes the code from the `/internal` and `/pkg` directories and nothing else.

### `/internal`
Private application and library code. This is the code you don't want others importing in their applications or libraries. 
Note that this layout pattern is enforced by the Go compiler itself. See the Go 1.4 release notes for more details. 

### `/pkg`
Library code that's ok to use by external applications (e.g., `/pkg/mypubliclib`). Other projects will import these libraries expecting them to work.
Note that the internal directory is a better way to ensure your private packages are not importable because it's enforced by Go. 
The `/pkg` directory is still a good way to explicitly communicate that the code in that directory is safe for use by others. 

### `/configs`
Configuration file templates or default configs.

### `/build`
Packaging and Continuous Integration related files.

### `/docs`
Design and user documents (in addition to your godoc generated documentation).

### `/bin`
Executable Binary location.

## Project Root Files

### `.gitignore`
List of git ignored files and folders.

### `go.mod`
List of Go Modules configurations such as module name, module list, go version and more.

## Godoc
Install godoc first (godoc isn't included in version `>=1.13`)
```
go install golang.org/x/tools/cmd/godoc@latest
```
Run it directly with `godoc -http:6060`. And open this [link](http://localhost:6060/pkg/jidniilman/quiz-master/?m=all)

## Golang CI Lint
Golang CI Lint is included with GitHub Actions.

But if you want to lint locally, Install Golang CI Lint first:
```
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin 

# test it
golangci-lint --version
```
Run lint with:
```
golangci-lint run
```
Check [https://golangci-lint.run/](https://golangci-lint.run/) for more commands and usages.
