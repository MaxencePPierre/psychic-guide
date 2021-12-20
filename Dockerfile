# syntax=docker/dockerfile:1

FROM golang:1.17

## Create an /app directory for source files
RUN mkdir /app
ADD . /app
WORKDIR /app

## Launch tests
RUN go test ./...

## Build the application
RUN go build -o main .

## Expose port
EXPOSE 4161

## Kick off the executable
CMD ["/app/main"]