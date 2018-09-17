package message

import (
	"context"

	"github.com/chatsex/domain"
	"github.com/chatsex/service"
	"github.com/go-kit/kit/endpoint"
)

type CreateRequest struct {
	UserNameSend    string `json:"user_name_send"`
	RoomNameReceive string `json:"room_name_receive"`
	Body            string `json:"body"`
}

type CreateResponse struct {
	MessageRoom *domain.MessageRoom `json:"message"`
}

func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		message := &domain.MessageRoom{UserNameSend: req.UserNameSend, RoomNameReceive: req.RoomNameReceive, Body: req.Body}
		res, err := s.MessageService.Create(ctx, message)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

type GetAllMessageRoomRequest struct {
	UserName string `json:"user_name"`
	RoomName string `json:"room_name"`
}

type GetAllMessageRoomResponse struct {
	Message []domain.MessageRoom `json:"messages"`
}

func MakeGetAllMessageRoomEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllMessageRoomRequest)
		res, err := s.MessageService.GetAllMessageRoom(ctx, req.RoomName, req.UserName)
		if err != nil {
			return nil, err
		}
		return GetAllMessageRoomResponse{Message: res}, nil
	}
}
