package domain

type Server struct {
	ConnectedUser       map[UUID]*User
	RoomCreated         map[UUID]*Room
	AddUserRoom         chan *UserRoom
	AddRoom             chan *Room
	NewInCommingMessage chan *MessageRoom
}
