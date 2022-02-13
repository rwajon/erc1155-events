FROM golang:1.17-alpine AS development
WORKDIR /app
COPY . ./
RUN apk add --update curl
RUN go mod download

FROM development AS builder
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM alpine:latest AS production
ENV GIN_MODE=release
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/app ./
CMD ["./app"]  
