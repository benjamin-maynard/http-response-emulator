# Define the base image, see https://hub.docker.com/_/golang for official Go images
FROM golang:1.13.10-alpine3.11

# Define the Working Directory, and copy in the source code from locally
WORKDIR /go/src/app
COPY ./http-response-emulator .

# Build the container
RUN go build main.go

# Run the service
CMD ["./main"]