package message

import (
	"context"

	"github.com/chatrealtime/domain"
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

func (s *pgService) Create(ctx context.Context, o *domain.MessageRoom) (*domain.MessageRoom, error) {
	user := &domain.User{}
	if err := s.db.Where("name = ?", o.UserNameSend).Find(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNameNotExist
		}
		return nil, err
	}

	room := &domain.Room{}
	if err := s.db.Where("name = ?", o.RoomNameReceive).Find(room).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRoomNameNotExist
		}
		return nil, err
	}

	userRoom := &domain.UserRoom{}
	if err := s.db.Where("user_id = ? AND room_id = ?", user.ID, room.ID).Find(userRoom).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotYetJoinRoom
		}
		return nil, err
	}

	message := &domain.MessageRoom{UserIDSend: &user.ID,
		UserNameSend:    user.Name,
		RoomIDReceive:   &room.ID,
		RoomNameReceive: room.Name,
		Body:            o.Body,
	}
	err := s.db.Create(message).Error
	if err != nil {
		return nil, err
	}
	return message, nil

}

func (s *pgService) GetAllMessageRoom(ctx context.Context, roomName string, userName string) ([]domain.MessageRoom, error) {
	room := &domain.Room{}
	if err := s.db.Where("name = ?", roomName).Find(room).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRoomNameNotExist
		}
		return nil, err
	}

	user := &domain.User{}
	if err := s.db.Where("name = ?", userName).Find(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNameNotExist
		}
		return nil, err
	}

	userRoom := &domain.UserRoom{}
	if err := s.db.Where("user_id = ? AND room_id = ?", user.ID, room.ID).Find(userRoom).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotYetJoinRoom
		}
		return nil, err
	}

	messageRoom := []domain.MessageRoom{}
	if err := s.db.Where("room_id_receive = ?", room.ID.String()).Find(&messageRoom).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return messageRoom, nil
		}
		return nil, err
	}
	return messageRoom, nil

}
