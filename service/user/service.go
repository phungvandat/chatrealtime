package user

import (
	"context"

	"github.com/chatsex/domain"
)

type Service interface {
	Register(ctx context.Context, o *domain.User) error
	Login(ctx context.Context, o *domain.User) (*domain.User, error)
	Get(ctx context.Context, userName string) (*domain.User, error)
}
