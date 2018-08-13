package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeBookEndpoint(svc BookingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(slot)
		br, err := svc.Book(req)
		if err != nil {
			return booking{br.Time, br.Court, err.Error()}, nil
		}
		return br, nil
	}
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

func parse() {
}

func submit() {
}
