package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger	log.Logger
	next	BookingService
}

func (mw loggingMiddleware) Book(b slot) (output booking, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "book",
			"days", b.Days,
			"court", b.Court,
			"hour", b.Hour,
			"min", b.Min,
			"timeslot", b.TimeSlot,
			"output", output.Time,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Book(b)
	return
}
