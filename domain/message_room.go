package domain

type MessageRoom struct {
	Model
	UserNameSend    string `json:"user_name_send"`
	UserIDSend      *UUID  `json:"user_id_send"`
	RoomNameReceive string `json:"room_name_receive"`
	RoomIDReceive   *UUID  `json:"room_id_receive"`
	Body            string `json:"body"`
}
