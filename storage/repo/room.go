package repo

import "time"
type RoomRepo struct{
	ID int64
	HotelId int64
	RoomNum int64
	BookedFrom time.Time
	BokedTo time.Time
}
type GetallRoomParams struct {
	Limit  int64
	Page   int64
	Search string
	SortBy string
}
type GetallRoomResponse struct {
	Users []*RoomRepo
	Count int64
}


type RoomStorageI interface {
	Create(Room *RoomRepo) (*RoomRepo, error)
	Get(id int64) (*RoomRepo, error)
	GetAll(params *GetallRoomParams) (*GetallRoomResponse, error)
	Update(Room *RoomRepo) (*RoomRepo, error)
	Delete(id int64) error
}