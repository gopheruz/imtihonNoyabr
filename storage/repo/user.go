package repo

import "time"

type UserRepo struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
	ImageUrl  string
	UserType  string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type UserStorageI interface {
	Create(User *UserRepo) (*UserRepo, error)
	Get(id int64) (*UserRepo, error)
	GetAll(params *GetallUsersParams) (*GetallUsersResponse, error)
	Update(User *UserRepo) (*UserRepo, error)
	Delete(id int64) error
}

type GetallUsersParams struct {
	Limit  int64
	Page   int64
	Search string
	Sortby string
}

type GetallUsersResponse struct {
	Users []*UserRepo
	Count int64
}
