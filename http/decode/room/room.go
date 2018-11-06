package room

import (
	"context"
	"encoding/json"
	"net/http"

	roomEndpoint "github.com/chatrealtime/endpoints/room"
)

func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req roomEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req roomEndpoint.FindRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func GetAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req roomEndpoint.GetAllRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
