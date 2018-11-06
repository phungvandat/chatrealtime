package service

import (
	"github.com/chatrealtime/service/message"
	"github.com/chatrealtime/service/room"
	"github.com/chatrealtime/service/user"
)

type Service struct {
	UserService    user.Service
	RoomService    room.Service
	MessageService message.Service
}
