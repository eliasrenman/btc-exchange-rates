FROM golang:alpine AS builder

WORKDIR /app

COPY src/go.mod src/go.sum ./

RUN go mod download

COPY ./src .

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest


# Copy the compiled binary from the previous stage
COPY --from=builder /app/app /app/app

# Set the working directory for the application
WORKDIR /app

# Command to run the executable
CMD ["./app"]
