package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc BookingService
	svc = bookingService{}
	svc = loggingMiddleware{logger, svc}

	bookHandler := httptransport.NewServer(
		makeBookEndpoint(svc),
		decodeBookRequest,
		encodeResponse,
	)

	http.Handle("/book", bookHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
