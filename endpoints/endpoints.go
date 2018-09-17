package endpoints

import (
	"github.com/chatsex/endpoints/message"
	"github.com/chatsex/endpoints/room"
	"github.com/chatsex/endpoints/user"
	"github.com/chatsex/service"
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
