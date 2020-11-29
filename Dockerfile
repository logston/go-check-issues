FROM golang:1.15 AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o app .

FROM alpine:latest
WORKDIR /
COPY --from=builder /build/app .
CMD ["/app"]
