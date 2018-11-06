package room

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

func (s *pgService) GetAllUser(ctx context.Context, roomName string, userName string) ([]domain.User, error) {
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

	lUser := []domain.User{}
	if err := s.db.Joins("join USER_ROOMS on USER_ROOMS.user_id = USERS.id").Where("USER_ROOMS.room_id = ?", room.ID).Find(&lUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return lUser, nil
		}
		return nil, err
	}
	return lUser, nil
}

func (s *pgService) GetAll(ctx context.Context, userName string) ([]domain.Room, error) {
	user := &domain.User{}
	if err := s.db.Where("name = ?", userName).Find(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNameNotExist
		}
		return nil, err
	}
	room := []domain.Room{}

	if err := s.db.Joins("join USER_ROOMS on USER_ROOMS.room_id = ROOMS.id").Where("USER_ROOMS.user_id = ?", user.ID).Find(&room).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return room, nil
		}
		return nil, err
	}
	return room, nil
}

func (s *pgService) Find(ctx context.Context, roomName string, userName string) ([]domain.MessageRoom, error) {
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

	err := s.db.Create(&domain.UserRoom{UserID: &user.ID, RoomID: &room.ID}).Error
	if err != nil {
		return nil, err
	}

	resMessage := []domain.MessageRoom{}
	err = s.db.Where("room_id_receive = ?", room.ID).Find(&resMessage).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return resMessage, nil
		}
		return nil, err
	}
	return resMessage, nil

}

func (s *pgService) Create(ctx context.Context, roomName string, userName string) error {
	if err := s.db.Where("name = ?", roomName).Find(&domain.Room{}).Error; err == nil {
		return ErrRoomNameExist
	}
	user := &domain.User{}
	if err := s.db.Where("name = ?", userName).Find(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrUserNameNotExist
		}
		return err
	}
	transaction := s.db.Begin()

	room := &domain.Room{Name: roomName}
	err := transaction.Create(room).Error

	if err != nil {
		return err
	}

	err = transaction.Create(&domain.UserRoom{UserID: &user.ID, RoomID: &room.ID}).Error
	if err != nil {
		transaction.Rollback()
		transaction.Commit()
		return err
	}
	transaction.Commit()
	return nil
}
