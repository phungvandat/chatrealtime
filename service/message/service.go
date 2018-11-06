package message

import (
	"context"

	"github.com/chatrealtime/domain"
)

type Service interface {
	GetAllMessageRoom(ctx context.Context, roomName string, userName string) ([]domain.MessageRoom, error)
	Create(ctx context.Context, o *domain.MessageRoom) (*domain.MessageRoom, error)
}
