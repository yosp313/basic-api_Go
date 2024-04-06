FROM golang:1.22.1

WORKDIR /usr/src

COPY go.mod go.sum ./
RUN go mod tidy && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

EXPOSE 8080

CMD ["app"]
