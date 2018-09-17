package service

import (
	"github.com/chatsex/service/message"
	"github.com/chatsex/service/room"
	"github.com/chatsex/service/user"
)

type Service struct {
	UserService    user.Service
	RoomService    room.Service
	MessageService message.Service
}
