package main

import (
	"errors"
)

type BookingService interface {
	Book(slot) (booking, error)
}

type bookingService struct{}

func (bookingService) Book(b slot) (booking, error) {
	// read environment variables defining tynemouth squash url
	// parse
	// submit
	if b.Min == "" {
		return booking{}, ErrParameter
	}
	br := booking{b.Hour + ":" + b.Min + " 21/08/2018", b.Court, ""}
	return br, nil
}

var ErrParameter = errors.New("Empty parameter")
