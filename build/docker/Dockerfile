FROM golang:latest
LABEL maintainer="Nika <veronika.kochugova@gmail.com>"
RUN go get github.com/akamensky/argparse
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
RUN go build -o ./build ./cmd/lab1