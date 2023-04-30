# GA(Gist Attendance) Backend Server

This is the backend server for the Skrr app.

## Getting Started

To get started with this project, you need to do the following:

1. Clone this repository.
2. Install the dependencies by running go mod tidy.
3. Start the server by running `go run main.go` or you can also build an docker image using `Dockerfile` on root folder.

!! Can't be executed in this github because private key is not exposed in public 

<br />
To run a test

```bash
go test -coverprofile cover.prof ./...
go tool cover -html=cover.prof
```
