package user

import (
	"context"

	"github.com/chatsex/domain"
	"github.com/jinzhu/gorm"
)

type pgService struct {
	db *gorm.DB
}

func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

func (s *pgService) Get(ctx context.Context, userName string) (*domain.User, error) {
	user := &domain.User{}
	if err := s.db.Where("name = ?", userName).Find(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNameNotExist
		}
		return nil, err
	}
	return user, nil
}

func (s *pgService) Register(ctx context.Context, o *domain.User) error {
	if err := s.db.Where("name = ?", o.Name).Find(&domain.User{}).Error; err == nil {
		return ErrUserNameExist
	}

	if err := s.db.Where("phone_number = ?", o.PhoneNumber).Find(&domain.User{}).Error; err == nil {
		return ErrPhoneNumberExist
	}

	return s.db.Create(o).Error
}

func (s *pgService) Login(ctx context.Context, o *domain.User) (*domain.User, error) {
	if err := s.db.Where("name = ? AND password = ?", o.Name, o.Password).Find(&domain.User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrorLoginInformation
		}
		return nil, err
	}
	return o, nil
}
