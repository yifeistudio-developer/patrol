FROM golang:latest AS builder

WORKDIR /usr/src/app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o payment ./cmd/main.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /usr/src/app/payment ./payment
CMD ["./payment"]
