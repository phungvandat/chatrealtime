package websoc

import (
	"context"
	"fmt"
	"log"

	"github.com/chatsex/domain"
	"github.com/gorilla/websocket"
)

const channelBufSizeUser = 1000

type UserConn struct {
	User            *domain.User
	Conn            *websocket.Conn
	WebSoc          *WebSoc
	OutGoingMessage chan *domain.MessageRoom
	DoneCh          chan bool
}

func NewUserConn(conn *websocket.Conn, webSoc *WebSoc, user *domain.User) (*UserConn, error) {
	if conn == nil {
		return nil, fmt.Errorf("conn required")
	}
	if webSoc == nil {
		return nil, fmt.Errorf("webSoc required")
	}
	if user == nil {
		return nil, fmt.Errorf("user required")
	}
	return &UserConn{
		User:            user,
		Conn:            conn,
		WebSoc:          webSoc,
		OutGoingMessage: make(chan *domain.MessageRoom, channelBufSizeUser),
		DoneCh:          make(chan bool),
	}, nil
}

//listen write and read message
func (u *UserConn) Listen() {
	go u.ListenWriteMessage()
	go u.ListenReadeMessage()

}

func (u *UserConn) ListenWriteMessage() {
	for {
		select {
		case msg := <-u.OutGoingMessage:
			u.Conn.WriteJSON(msg)
		case <-u.DoneCh:
			u.WebSoc.DisconnectUser <- u
			u.DoneCh <- true
			return
		}
	}
}

func (u *UserConn) ListenReadeMessage() {
	for {
		select {
		case <-u.DoneCh:
			u.WebSoc.DisconnectUser <- u
			u.DoneCh <- true
			return
		default:
			var msg domain.MessageRoom
			err := u.Conn.ReadJSON(&msg)
			if err != nil {
				u.DoneCh <- true
				u.WebSoc.ErrChannel <- err
				return
			} else {
				message, err := u.WebSoc.Service.MessageService.Create(context.Background(), &msg)
				if err != nil {
					log.Println("failed get user by error ", err)
				}
				u.WebSoc.NewInComingMessage <- message
			}
		}
	}
}

func (u *UserConn) Write(msg *domain.MessageRoom) {
	select {
	case u.OutGoingMessage <- msg:
	default:
		u.WebSoc.DisconnectUser <- u
		u.WebSoc.ErrChannel <- fmt.Errorf("disconnect user: ", u.User.Name)
	}
}
