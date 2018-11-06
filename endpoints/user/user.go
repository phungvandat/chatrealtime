package user

import (
	"context"

	"github.com/chatrealtime/domain"
	"github.com/chatrealtime/service"
	"github.com/go-kit/kit/endpoint"
)

//Register
type RegisterData struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterRequest struct {
	UserData RegisterData `json:"user"`
}

type RegisterResponse struct {
	Status bool `json:"status"`
}

func MakeRegisterEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		user := &domain.User{
			Name:        req.UserData.Name,
			Password:    req.UserData.Password,
			PhoneNumber: req.UserData.PhoneNumber,
		}
		err := s.UserService.Register(ctx, user)
		if err != nil {
			return nil, err
		}
		return RegisterResponse{Status: true}, nil
	}
}

//Login
type LoginData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRequest struct {
	UserData LoginData `json:"user"`
}

type LoginResponse struct {
	User domain.User `json:"user"`
}

func MakeLoginEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		user := &domain.User{Name: req.UserData.Name, Password: req.UserData.Password}
		res, err := s.UserService.Login(ctx, user)
		if err != nil {
			return nil, err
		}
		return LoginResponse{User: *res}, nil
	}
}
