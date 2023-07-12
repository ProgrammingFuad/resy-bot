# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./

# Add all the Go files in the src directory to the container's working directory
ADD src/*.go /app/

RUN go mod download 

# Build the Go files
RUN go build -o app *.go

# Run the built application when the container starts
CMD [ "./app" ]

