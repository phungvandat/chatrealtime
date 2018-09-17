package message

import "net/http"

var (
	ErrUserNameNotExist   = errUserNameNotExist{}
	ErrRoomNameNotExist   = errRoomNameNotExist{}
	ErrUserNotYetJoinRoom = errUserNotYetJoinRoom{}
)

type errUserNameNotExist struct{}

func (errUserNameNotExist) Error() string {
	return "user name not exist"
}

func (errUserNameNotExist) StatusCode() int {
	return http.StatusBadRequest
}

type errRoomNameNotExist struct{}

func (errRoomNameNotExist) Error() string {
	return "room name not exist"
}

func (errRoomNameNotExist) StatusCode() int {
	return http.StatusBadRequest
}

type errUserNotYetJoinRoom struct{}

func (errUserNotYetJoinRoom) Error() string {
	return "user not yet join room"
}

func (errUserNotYetJoinRoom) StatusCode() int {
	return http.StatusBadRequest
}
