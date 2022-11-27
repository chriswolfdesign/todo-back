FROM golang:latest

WORKDIR /app

COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
COPY cmd /app/cmd
COPY db /app/db
COPY endpoints /app/endpoints
COPY model /app/model

RUN go mod download

RUN go build -o app cmd/cmd.go

CMD ["./app"]