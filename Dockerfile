# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /prevention_productivity

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o ./prevention_productivity

EXPOSE 8080

CMD ["/prevention_productivity"]