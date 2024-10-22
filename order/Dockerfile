FROM golang:latest AS builder

WORKDIR /usr/src/app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o order ./cmd/main.go

FROM scratch
COPY --from=builder /usr/src/app/order ./order
CMD ["./order"]

