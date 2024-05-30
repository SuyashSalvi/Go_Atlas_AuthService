# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:1.16 as builder

# Copy all files from the current directory to the /authservice2 directory.
COPY . /authservice2

# Set the current working directory inside the container to /authservice2.
WORKDIR /authservice2

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Build the Go app. Output a binary named authservice2.
RUN go build -o authservice2 .

# Start a new stage from scratch.
FROM debian:bullseye

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /authservice2/authservice2 /usr/local/bin/authservice2

# Run the binary program.
ENTRYPOINT ["/usr/local/bin/authservice2"]
