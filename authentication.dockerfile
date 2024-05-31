# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

# Build the authApp binary using the go tool
RUN CGO_ENABLED=0 go build -o authApp ./cmd/api

# Make the authApp binary executable
RUN chmod +x /app/authApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# Copy the authApp binary from the builder stage to the base image
COPY --from=builder /app/authApp /app

CMD [ "/app/authApp"]
