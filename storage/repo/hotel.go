package repo

type HotelRepo struct {
	ID          int64
	HotelName   string
	HotelAdress string
	Description string
	Images      []*HotelImages
}

type HotelImages struct {
	ID       int64
	HotelID  int64
	ImageUrl string
}

type HotelStorageI interface {
	Create(Hotel *HotelRepo) (*HotelRepo, error)
	Get(id int64) (*HotelRepo, error)
	GetAll(params *GetallHotelParams) (*GetallHotelResponse, error)
	Update(Hotel *HotelRepo) (*HotelRepo, error)
	Delete(id int64) error
}
type GetallHotelParams struct {
	Limit  int64
	Page   int64
	Search string
	SortBy string
}
type GetallHotelResponse struct {
	Hotels []*HotelRepo
	Count  int64
}
