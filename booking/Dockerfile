FROM golang:1.10.3 as builder
WORKDIR /go/src/github.com/xzx1kf/squash/booking
RUN go get -d -v github.com/prometheus/client_golang/prometheus
RUN go get -d -v github.com/go-kit/kit/endpoint
RUN go get -d -v github.com/go-kit/kit/log
RUN go get -d -v github.com/go-kit/kit/metrics
RUN go get -d -v github.com/go-kit/kit/transport/http
COPY main.go .
COPY transport.go .
COPY service.go .
COPY instrumenting.go .
COPY logging.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o booking .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/xzx1kf/squash/booking .
CMD ["./booking"]
