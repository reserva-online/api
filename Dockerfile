FROM golang:1.16-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/schedule-api

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/schedule-api ./cmd

# This container exposes port 8000 to the outside world
EXPOSE 8000

# Run the binary program produced by `go install`
CMD ["./out/schedule-api"]