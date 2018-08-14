package main

import (
	"net/http"
	"os"

    stdprometheus "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
    kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

    fieldKeys := []string{"method", "error"}
    requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
        Namespace: "my_group",
        Subsystem: "booking_service",
        Name:      "request_count",
        Help:      "Number of requests received.",
    }, fieldKeys)
    requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
        Namespace: "my_group",
        Subsystem: "booking_service",
        Name:      "request_latency_microseconds",
        Help:      "Total duration of requests in microseconds.",
    }, fieldKeys)

	var svc BookingService
	svc = bookingService{}
	svc = loggingMiddleware{logger, svc}
    svc = instrumentingMiddleware{requestCount, requestLatency, svc}

	bookHandler := httptransport.NewServer(
		makeBookEndpoint(svc),
		decodeBookRequest,
		encodeResponse,
	)

	http.Handle("/book", bookHandler)
    http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
