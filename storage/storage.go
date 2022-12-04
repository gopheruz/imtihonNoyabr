package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/nurmuhammaddeveloper/imtihonNoyabr/storage/postgres"
	"github.com/nurmuhammaddeveloper/imtihonNoyabr/storage/repo"
)


type StorageI interface{
	Hotel() repo.HotelStorageI
	User() repo.UserStorageI
	HotelOrder() repo.HotelOrderStorageI
	Room() repo.RoomStorageI
}
type StoragePg struct {
	hotelRepo   repo.HotelStorageI
	userRepo    repo.UserStorageI
	hotelOrders repo.HotelOrderStorageI
	roomRepo repo.RoomStorageI
}



func NewStoragePg(db *sqlx.DB)StorageI{
	return &StoragePg{
		hotelRepo: postgres.NewHotel(db),
		userRepo: postgres.NewUser(db),
		hotelOrders: postgres.NewOrder(db),
	}
}

func (s *StoragePg) Hotel() repo.HotelStorageI{
	return s.hotelRepo
}
func (s *StoragePg) User() repo.UserStorageI{
	return s.userRepo
}
func (s *StoragePg) HotelOrder() repo.HotelOrderStorageI{
	return s.hotelOrders
}

func (s *StoragePg)Room()repo.RoomStorageI{
	return s.roomRepo
}
