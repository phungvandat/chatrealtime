package domain

type UserRoom struct {
	Model
	UserID *UUID `json:"user_id"`
	RoomID *UUID `json:"room_id"`
}
