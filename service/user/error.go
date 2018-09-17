package user

import "net/http"

var (
	ErrUserNameExist      = errUserNameExist{}
	ErrUserNameNotExist   = errUserNameNotExist{}
	ErrPhoneNumberExist   = errPhoneNumberExist{}
	ErrorLoginInformation = errorLoginInformation{}
)

type errUserNameExist struct{}

func (errUserNameExist) Error() string {
	return "user name exist"
}

func (errUserNameExist) StatusCode() int {
	return http.StatusBadRequest
}

type errUserNameNotExist struct{}

func (errUserNameNotExist) Error() string {
	return "user name not exist"
}

func (errUserNameNotExist) StatusCode() int {
	return http.StatusBadRequest
}

type errPhoneNumberExist struct{}

func (errPhoneNumberExist) Error() string {
	return "phone number exist"
}

func (errPhoneNumberExist) StatusCode() int {
	return http.StatusBadRequest
}

type errorLoginInformation struct{}

func (errorLoginInformation) Error() string {
	return "Incorrect login information"
}

func (errorLoginInformation) StatusCode() int {
	return http.StatusBadRequest
}
