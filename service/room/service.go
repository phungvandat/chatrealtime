package room

import (
	"context"

	"github.com/chatrealtime/domain"
)

type Service interface {
	Create(ctx context.Context, roomName string, userName string) error
	Find(ctx context.Context, roomName string, userName string) ([]domain.MessageRoom, error)
	GetAll(ctx context.Context, userName string) ([]domain.Room, error)
	GetAllUser(ctx context.Context, roomName string, userName string) ([]domain.User, error)
}
