package message

import (
	"context"
	"encoding/json"
	"net/http"

	messageEndpoint "github.com/chatsex/endpoints/message"
)

func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req messageEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func GetAllMessageRoomRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req messageEndpoint.GetAllMessageRoomRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
