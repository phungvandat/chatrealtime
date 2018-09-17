package room

import (
	"context"

	"github.com/chatsex/domain"
	"github.com/chatsex/service"
	"github.com/go-kit/kit/endpoint"
)

//Create
type RoomData struct {
	UserName string `json:"user_name"`
	RoomName string `json:"room_name"`
}

type CreateRequest struct {
	Room RoomData `json:"room"`
}

type CreateResponse struct {
	Status bool `json:"status"`
}

func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		err := s.RoomService.Create(ctx, req.Room.RoomName, req.Room.UserName)
		if err != nil {
			return nil, err
		}
		return CreateResponse{Status: true}, nil
	}
}

type FindRequest struct {
	Room RoomData `json:"room"`
}

type FindResponse struct {
	Message []domain.MessageRoom `json:"messages"`
}

func MakeFindEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FindRequest)
		res, err := s.RoomService.Find(ctx, req.Room.RoomName, req.Room.UserName)
		if err != nil {
			return nil, err
		}
		return FindResponse{Message: res}, nil
	}
}

type GetAllRequest struct {
	UserName string `json:"user_name"`
}

type GetALlResponse struct {
	Room []domain.Room `json:"rooms"`
}

func MakeGetAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllRequest)
		res, err := s.RoomService.GetAll(ctx, req.UserName)
		if err != nil {
			return nil, err
		}
		return GetALlResponse{Room: res}, nil
	}
}
