# Quiz-Master
[![wakatime](https://wakatime.com/badge/user/90095cec-bff0-4e71-9e19-19fbac89fb11/project/9a69d0dd-130c-49b6-93ab-2a629bd2b36b.svg)](https://wakatime.com/badge/user/90095cec-bff0-4e71-9e19-19fbac89fb11/project/9a69d0dd-130c-49b6-93ab-2a629bd2b36b)
![Golang CI Lint](https://github.com/jidniilman/Quiz-Master/actions/workflows/golangci-lint.yml/badge.svg)
[![go-ci](https://github.com/jidniilman/Quiz-Master/actions/workflows/go-ci.yml/badge.svg)](https://github.com/jidniilman/Quiz-Master/actions/workflows/go-ci.yml)
----
Quiz Master is Quipper’s Technical Exam for Backend Engineer - Web

## Overview
- This project is tracked by WakaTime
- This project built with Go 1.19
- This project created with Linux (Ubuntu WSL2 on Windows)
- This project developed and coded with Windows 10
- This project is tested on both Ubuntu and Windows 10
- Detailed time elapsed to write this project can be found [here (WakaTime)](https://wakatime.com/@deraven/projects/qstzzolwiz?start=2022-08-02&end=2022-08-08) 

## Checklist
- [x] App Functionality
- [x] Push to Github for first time
- [x] Cleanup Readme
- [x] Write Unit Tests
- [x] Write UNIX Executable Script `/bin/setup` and `/bin/quiz_master`
- [x] Implement GitHub Actions
- [x] Write proper documentation
- [x] Implement Golang CI Lint
- [x] Write Benchmark
- [x] Provide Dockerfile
- [x] Provide Possible Improvement

## Quickstart
Run this to download dependencies, compile, and run our unit tests:
```
bin/setup
```
Run this to execute our program (it's basically `go run ./...`):
```
bin/quiz_master
```
You can also run from cmd entrypoint at `/cmd/quiz_master` and run this:
```
go run main.go
```

## How to Use
After you run the program, you can input commands and arguments below:

| Short | Command           | Arguments                          | Description                                |
|-------|-------------------|------------------------------------|--------------------------------------------|
| `cq`  | `create_question` | `<number>` `<question>` `<answer>` | Create a question with a number and answer |             
| `uq`  | `update_question` | `<number>` `<question>` `<answer>` | Update a question with a number and answer |                                            
| `dq`  | `delete_question` | `<number>`                         | Delete a question by given number          |                                            
| `q`   | `question`        | `<number>`                         | Show a question by given number            |                                            
| `qs`  | `questions`       |                                    | Show all questions                         |                                            
| `aq`  | `answer_question` | `<number>` `<answer>`              | Answer a question by given number          |                                            
|       | `version`         |                                    | Show app version                           |                                            
|       | `help`            |                                    | Show help                                  |                                            
|       | `exit`            |                                    | Close the program                          |                                            

Example 1:
```
create_question 1 "How many letters are there in the English alphabet?" 26
```
Example 2:
```
create_question 2 “How many vowels are there in the English alphabet?” 5544879599
```
Example 3:
```
answer_question 2 "five billion five hundred forty four million eight hundred seventy nine thousand five hundred ninety nine"
```

## Library Used
We use our own built solution as much as possible. Everything is self-made except for go standard package and:
- [Testify](https://github.com/stretchr/testify) for Unit Testing purpose

## Coding Convention
Golang is strict by default. But we are using some guidelines from :
- [Effective Go](https://go.dev/doc/effective_go)
- [CockroachDB Engineering Standards](https://wiki.crdb.io/wiki/spaces/CRDB/pages/181371303/Go+Golang+coding+guidelines)
- [Golang CI Lint](https://golangci-lint.run/)
- GoLand IDE formatting

## Unit Tests
Mainly for our unit testing, we are using table driven test approach. But we are using little sample to do our tests.

Run all unit tests with 
```
go test ./...
```
or 
```
# with verbose output.
go test -v ./...
```
You can run individual unit test with 
```
go test -v <package-name>
```
Package Name (main entrypoint are excluded):
```
jidniilman/quiz-master/internal/cli      # Handling CLI application and execution
jidniilman/quiz-master/internal/question # Our question model
jidniilman/quiz-master/pkg/command       # Our data communication protocol
jidniilman/quiz-master/pkg/utils         # Some useful helpers
```
To see the full package list, use 
```
go list ./...
```

## Benchmark
Only `/pkg/command` and `/pkg/utils` that have benchmark. We can do benchmark for `/internal` but it is not necessary.

Run benchmark with:
```
go test -v ./... -bench=. -benchmem
```

## Project Layout/Structure
In this project, I'm using the layout standards from [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
I remove some of unused folders such as: `/api`, `/web`, `/deployments`, `/test`, `/configs`, `/docs`, and `/build`

### `/cmd`
Main applications for this project. The directory name for each application match the name of the executable (e.g., `/cmd/quiz_master`).
It's common to have a small main function that imports and invokes the code from the `/internal` and `/pkg` directories and nothing else.

### `/internal`
Private application and library code. This is the code you don't want others importing in their applications or libraries. 
Note that this layout pattern is enforced by the Go compiler itself. See the Go 1.4 release notes for more details. 

### `/pkg`
Library code that's ok to use by external applications (e.g., `/pkg/utils`). Other projects will import these libraries expecting them to work.
Note that the internal directory is a better way to ensure your private packages are not importable because it's enforced by Go. 
The `/pkg` directory is still a good way to explicitly communicate that the code in that directory is safe for use by others. 

### `/bin`
Executable UNIX Scripts and Binary location.

## Project Root Files

### `.gitignore`
List of git ignored files and folders.

### `go.mod`
List of Go Modules configurations such as module name, module list, go version and more.

### `.github/workflows/*.yml`
List of GitHub Actions workflow we are using.

## Godoc
Install godoc first (godoc isn't included in version `>=1.13`)
```
go install golang.org/x/tools/cmd/godoc@latest
```

### Serve via HTTP
Run it directly with `godoc -http localhost:6060`. And open this [link](http://localhost:6060/pkg/jidniilman/quiz-master/?m=all)

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

## Docker
We provide two-stage build docker image for you.

Build docker with:
```
docker build -t jidniilman/quiz-master .
```
See created docker images with:
```
docker images
```
Run the docker image that we previously created:
```
docker run -it jidniilman/quiz-master:latest
```

## Possible Improvement
- Make benchmark for `/intenal`
- Improve performance of each function
- Separate into microservice architecture for interesting training with gRPC/Protobuf or other Pub/Sub brokers.
- Saga Pattern is interesting to learn
- API Gateway Pattern with Kong is also interesting