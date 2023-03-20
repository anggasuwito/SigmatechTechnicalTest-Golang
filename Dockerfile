FROM golang:1.18-alpine AS builder
WORKDIR /app/sigmatech-golang
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o sigmatech-golang.exe

FROM alpine
RUN apk update && apk add --no-cache tzdata
WORKDIR /app
COPY --from=builder /app/sigmatech-golang /app/
ENTRYPOINT ["./sigmatech-golang.exe"]