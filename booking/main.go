package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type BookingService interface {
	Book(context.Context, slot) (booking, error)
}

type bookingService struct{}

func (bookingService) Book(_ context.Context, b slot) (booking, error) {
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

type slot struct {
	Days     string `json:"days"`
	Court    string `json:"court"`
	Hour     string `json:"hour"`
	Min      string `json:"min"`
	TimeSlot string `json:"timeslot"`
}

type booking struct {
	Time  string `json:"time"`
	Court string `json:"court"`
	Err   string `json:"err,omitempty"`
}

func makeBookEndpoint(svc bookingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(slot)
		br, err := svc.Book(ctx, req)
		if err != nil {
			return booking{br.Time, br.Court, err.Error()}, nil
		}
		return br, nil
	}
}

func main() {
	svc := bookingService{}

	bookHandler := httptransport.NewServer(
		makeBookEndpoint(svc),
		decodeBookRequest,
		encodeResponse,
	)

	http.Handle("/book", bookHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request slot
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func parse() {
}

func submit() {
}
