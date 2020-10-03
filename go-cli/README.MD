## How to Build CLI application in golang
This is an sample application for building the CLI application in golang.
Please follow the [post](https://blog.nitishkumarsingh.xyz/2020-10-03-build-console-application-in-golang/) to understand how to build an CLI app.

### Create an golang app. Replace module name with your module name
```
mkdir go-cli
go mod init github.com/nitishkumar71/blog/go-cli
```
### Install [Cobra](https://github.com/spf13/cobra) library
```
go get -u github.com/spf13/cobra/cobra
```
### run application while development
```
go run main.go <<sub-command>> <<name of file/directory>> <<flag1>> <<flag2>>
```
### build CLI application
```
go build -o bin/go-cli
```