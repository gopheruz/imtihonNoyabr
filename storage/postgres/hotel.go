package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/nurmuhammaddeveloper/imtihonNoyabr/storage/repo"
)


type hotelRepo struct{
	db *sqlx.DB
}

func NewHotel(db *sqlx.DB) repo.HotelStorageI{
	return &hotelRepo{
		db :db,
	}
}

func (hoteldb *hotelRepo)Create(Hotel *repo.HotelRepo)(*repo.HotelRepo, error){
	return nil, nil
}
func (hoteldb *hotelRepo)Get(id int64)(*repo.HotelRepo, error){
	return nil, nil
}

func (hoteldb *hotelRepo)GetAll(params *repo.GetallHotelParams)(*repo.GetallHotelResponse, error){
	return nil, nil
}

func (hoteldb *hotelRepo)Update(Hotel *repo.HotelRepo)(*repo.HotelRepo, error){
	return nil, nil
}

func (hoteldb *hotelRepo)Delete(id int64)(error){
	return nil
}