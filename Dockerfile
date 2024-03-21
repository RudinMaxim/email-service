FROM golang:latest

WORKDIR /usr/service/email

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["go", "run", "cmd/main.go"]