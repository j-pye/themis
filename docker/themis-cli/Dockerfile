FROM golang:1.15.8

WORKDIR /app
COPY . .

RUN go get -d ./...
RUN go install ./...
RUN go build cmd/themis-cli/*.go