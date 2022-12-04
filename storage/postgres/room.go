package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/nurmuhammaddeveloper/imtihonNoyabr/storage/repo"
)

type roomRepo struct {
	db *sqlx.DB
}

func NewRoom(db *sqlx.DB) repo.RoomStorageI{
	return &roomRepo{
		db: db,
	}
}

func(roomdb *roomRepo) Create(Room *repo.RoomRepo)(*repo.RoomRepo, error){
	return nil, nil
}

func (roomdb *roomRepo)Get(id int64)(*repo.RoomRepo, error){
	return nil, nil
}

func(roomdb *roomRepo)GetAll(params *repo.GetallRoomParams)(*repo.GetallRoomResponse, error){
	return nil, nil
}

func (roomdb *roomRepo)Update(Room *repo.RoomRepo)(*repo.RoomRepo, error){
	return nil, nil
}

func (roomdb *roomRepo) Delete(id int64) error{
	return nil
}