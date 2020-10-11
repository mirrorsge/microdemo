FROM golang:1.13 as builder

WORKDIR /app

COPY . /app

RUN  GOPROXY="https://goproxy.io" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./gateway/main.go

FROM alpine as product

WORKDIR /app

RUN apk --no-cache add git ca-certificates

COPY --from=builder /app/main /app

CMD ["./main"]

EXPOSE 8080:8080

