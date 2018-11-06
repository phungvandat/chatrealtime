package user

import (
	"context"

	"github.com/chatrealtime/domain"
)

type Service interface {
	Register(ctx context.Context, o *domain.User) error
	Login(ctx context.Context, o *domain.User) (*domain.User, error)
	Get(ctx context.Context, userName string) (*domain.User, error)
}
