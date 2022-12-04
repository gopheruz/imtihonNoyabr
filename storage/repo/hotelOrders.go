package repo

import "time"

type HotelOrder struct {
	ID              int64
	HotelID         int64
	UserID          int64
	HotelR00mNumber int64
	BookedFrom time.Time
	BookedTo time.Time
}

type HotelOrderStorageI interface {
	Create(Order *HotelOrder) (*HotelOrder, error)
	Get(id int64) (*HotelOrder, error)
	GetAll(params *GetallHotelOrderParams) (*GetallHotelOrderResponse, error)
	Update(Order *HotelOrder) (*HotelOrder, error)
	Delete(id int64) error
}
type GetallHotelOrderParams struct {
	Limit  int64
	Page   int64
	Search string
	SortBy string
}
type GetallHotelOrderResponse struct {
	Orders []*HotelOrder
	Count int64
}