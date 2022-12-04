package postgres
	
import (
	"github.com/jmoiron/sqlx"
	"github.com/nurmuhammaddeveloper/imtihonNoyabr/storage/repo"
)

type orderRepo struct {
	db *sqlx.DB
}

func NewOrder(db *sqlx.DB) repo.HotelOrderStorageI {
	return &orderRepo{db: db}
}

func (orderdb *orderRepo) Create(Order *repo.HotelOrder) (*repo.HotelOrder, error) {
	return nil, nil
}
func (orderdb *orderRepo) Get(id int64) (*repo.HotelOrder, error) {
	return nil, nil
}
func (orderdb *orderRepo) GetAll(params *repo.GetallHotelOrderParams) (*repo.GetallHotelOrderResponse, error) {
	return nil, nil
}
func (orderdb *orderRepo) Update(Order *repo.HotelOrder) (*repo.HotelOrder, error) {
	return nil, nil
}
func (orderdb *orderRepo) Delete(id int64) error {
	return nil
}
