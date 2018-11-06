package websoc

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/chatrealtime/domain"
	"github.com/chatrealtime/service"
	"github.com/gorilla/websocket"
)

type WebSoc struct {
	Service            *service.Service
	ConnectedUser      map[string]*UserConn
	DisconnectUser     chan *UserConn
	NewInComingMessage chan *domain.MessageRoom
	ErrChannel         chan error
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}

func NewWebSoc(service *service.Service) (*WebSoc, error) {
	if service == nil {
		return nil, fmt.Errorf("service empty")
	}
	connectedUser := make(map[string]*UserConn)
	disconnectUser := make(chan *UserConn)
	newInComingMessage := make(chan *domain.MessageRoom, channelBufSizeUser)
	errChannel := make(chan error)
	return &WebSoc{
		Service:            service,
		ConnectedUser:      connectedUser,
		DisconnectUser:     disconnectUser,
		NewInComingMessage: newInComingMessage,
		ErrChannel:         errChannel,
	}, nil
}

func (ws *WebSoc) HandlerChat(w http.ResponseWriter, req *http.Request) {
	//upgrade from http to websocket
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("failed create connection websocket by error ", err)
		return
	}
	log.Println("upgrade from http to websocket success")

	//read message
	var message domain.MessageRoom
	err = conn.ReadJSON(&message)
	if err != nil {
		log.Println("failed reade body message by error ", err)
		return
	}
	log.Println("read message success")

	//get user from database
	user, err := ws.Service.UserService.Get(context.Background(), message.UserNameSend)
	if err != nil {
		log.Println("failed get user by error ", err)
		return
	}
	log.Println("get user from database success")

	//New user connection
	userConn, err := NewUserConn(conn, ws, user)
	if err != nil {
		log.Println("failed create userConn by error ", err)
		return
	}
	log.Println("New user connection success")

	//add user connection to user connected
	ws.ConnectedUser[user.Name] = userConn
	log.Println("add user connection to user connected")

	//save message to database
	messageRoom, err := ws.Service.MessageService.Create(context.Background(), &message)
	if err != nil {
		log.Println("failed create message by error ", err)
		return
	}
	log.Println("save message to database")

	//add message to channel new in coming message
	ws.NewInComingMessage <- messageRoom
	userConn.Listen()
}

func (ws *WebSoc) SendMessage(msg *domain.MessageRoom) {
	lUserSend, err := ws.Service.RoomService.GetAllUser(context.Background(), msg.RoomNameReceive, msg.UserNameSend)
	if err != nil {
		log.Println("failed get user to send message by error ", err)
		return
	}
	for _, u := range lUserSend {
		if ws.ConnectedUser[u.Name] != nil {
			ws.ConnectedUser[u.Name].Write(msg)
		}
	}
}
