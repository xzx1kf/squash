package main

import (
    "fmt"
    "time"

    "github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
    requestCount    metrics.Counter
    requestLatency  metrics.Histogram
    next            BookingService
}

func (mw instrumentingMiddleware) Book(b slot) (output booking, err error) {
    defer func(begin time.Time) {
        lvs := []string{"method", "book", "error", fmt.Sprint(err != nil)}
        mw.requestCount.With(lvs...).Add(1)
        mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
    }(time.Now())

    output, err = mw.next.Book(b)
    return
}
