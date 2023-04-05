FROM golang:1.20.0

WORKDIR /usr/src/app

# Hot reloading
RUN go install github.com/cosmtrek/air@latest

# copy files to container, then install, cleanup dependencies
COPY . .
RUN go mod tidy
