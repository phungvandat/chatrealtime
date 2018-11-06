package endpoints

import (
	"github.com/chatrealtime/endpoints/message"
	"github.com/chatrealtime/endpoints/room"
	"github.com/chatrealtime/endpoints/user"
	"github.com/chatrealtime/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	RegisterUser endpoint.Endpoint
	LoginUser    endpoint.Endpoint

	CreateRoom endpoint.Endpoint
	FindRoom   endpoint.Endpoint
	GetAllRoom endpoint.Endpoint

	GetAllMessageRoom endpoint.Endpoint
	CreateMessage     endpoint.Endpoint
}

func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		RegisterUser: user.MakeRegisterEndpoint(s),
		LoginUser:    user.MakeLoginEndpoint(s),

		CreateRoom: room.MakeCreateEndpoint(s),
		FindRoom:   room.MakeFindEndpoint(s),
		GetAllRoom: room.MakeGetAllEndpoint(s),

		GetAllMessageRoom: message.MakeGetAllMessageRoomEndpoint(s),
		CreateMessage:     message.MakeCreateEndpoint(s),
	}
}
