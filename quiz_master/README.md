# Quiz-Master
Quiz Master is Quipper’s Technical Exam for Backend Engineer - Web

## Overview
- This project is tracked by WakaTime
- This project built with Go 1.19
- This project init by Linux (Ubuntu WSL2 on Windows)
- This project developed and coded with Windows 10

## Todo
- [x] App Functionality
- [x] Push to Github for first time
- [ ] Cleanup Readme
- [ ] Write Unit Tests
- [ ] Write UNIX Executable Script `/bin/setup` and `/bin/quiz_master`
- [ ] Implement GitHub Actions
- [ ] Write proper documentation
- [ ] Implement Golang CI Lint
- [ ] Write Benchmark

## Tech used
- Go Modules
- Makefile (todo)

## Library Used
Use own logic as much as possible

## Coding Convention
TBD

## Unit Tests

## Benchmark

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