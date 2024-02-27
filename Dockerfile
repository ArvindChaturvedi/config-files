# Build stage
FROM golang:1.16 AS build
WORKDIR /app
COPY . .
RUN go build -o sample-app

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /app/sample-app .
CMD ["./sample-app"]