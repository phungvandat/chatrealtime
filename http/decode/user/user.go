package user

import (
	"context"
	"encoding/json"
	"net/http"

	userEndpoint "github.com/chatrealtime/endpoints/user"
)

func RegisterRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req userEndpoint.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func LoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req userEndpoint.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
