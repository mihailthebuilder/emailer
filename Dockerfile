FROM golang:1.20

EXPOSE 8080/tcp

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go ./

ENV GIN_MODE=release
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]