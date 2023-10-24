FROM golang:latest

COPY . ./app
WORKDIR ./app

CMD ["go", "run", "cmd/cmd.go"]
