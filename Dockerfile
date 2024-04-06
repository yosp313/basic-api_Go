FROM golang:latest

WORKDIR /usr/src

COPY go.mod go.sum .
RUN go mod tidy && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

CMD ["app"]
