FROM golang:1.23-alpine AS builder
WORKDIR /go/src/github.com/rubnsbarbosa/dhl
COPY . .
RUN go get && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dhl .

FROM alpine:latest
RUN apk --update add ca-certificates
COPY --from=builder /go/src/github.com/rubnsbarbosa/dhl/dhl .
ENTRYPOINT ["./dhl"]
