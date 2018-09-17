package room

import "net/http"

var (
	ErrRoomNameExist      = errRoomNameExist{}
	ErrUserNameNotExist   = errUserNameNotExist{}
	ErrRoomNameNotExist   = errRoomNameNotExist{}
	ErrUserNotYetJoinRoom = errUserNotYetJoinRoom{}
)

type errRoomNameExist struct{}

func (errRoomNameExist) Error() string {
	return "room name exist"
}

func (errRoomNameExist) StatusCode() int {
	return http.StatusBadRequest
}

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
