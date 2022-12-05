package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/nurmuhammaddeveloper/imtihonNoyabr/pkg/utils"
	"github.com/nurmuhammaddeveloper/imtihonNoyabr/storage/querys"
	"github.com/nurmuhammaddeveloper/imtihonNoyabr/storage/repo"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{db: db}
}
func (usrdb *userRepo) Create(User *repo.UserRepo) (*repo.UserRepo, error) {
	hashedpw, err := utils.HashPassword(User.Password)
	if err != nil {
		return nil, err
	}
	if User.UserType == "" {
		User.UserType = "user"
	}
	rows := usrdb.db.QueryRow(querys.UserCreateQuery,
		User.FirstName,
		User.LastName,
		User.Email,
		hashedpw,
		User.UserType,
		User.ImageUrl,
	)
	err = rows.Scan(
		&User.ID,
		&User.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return User, err
}
func (usrdb *userRepo) Get(id int64) (*repo.UserRepo, error) {
	var GettedUser repo.UserRepo
	row := usrdb.db.QueryRow(querys.UseGetQuery, id)
	err := row.Scan(
		&GettedUser.FirstName,
		&GettedUser.LastName,
		&GettedUser.Email,
		&GettedUser.ImageUrl,
		&GettedUser.UserType,
		&GettedUser.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	GettedUser.ID = id
	return &GettedUser, nil
}
func (usrdb *userRepo) GetAll(params *repo.GetallUsersParams) (*repo.GetallUsersResponse, error) {
	result := repo.GetallUsersResponse{
		Users: make([]*repo.UserRepo, 0),
	}
	offset := (params.Page - 1) * params.Limit
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)
	filter := "WHERE deleted_at is null"
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(` 
			and first_name ILIKE '%s' OR last_name ILIKE '%s' OR email ILIKE '%s'`,
			str, str, str)
	}

	query := querys.GetAllUserQuery + filter + " ORDER BY created_at " + params.Sortby + " " + limit
	rows, err := usrdb.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var GettedUser repo.UserRepo
		err = rows.Scan(
			&GettedUser.ID,
			&GettedUser.FirstName,
			&GettedUser.LastName,
			&GettedUser.Email,
			&GettedUser.Password,
			&GettedUser.ImageUrl,
			&GettedUser.UserType,
			&GettedUser.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result.Users = append(result.Users, &GettedUser)
	}
	queryCount := `SELECT count(1) FROM users ` + filter
	err = usrdb.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (usrdb *userRepo) Update(User *repo.UserRepo) (*repo.UserRepo, error) {
	rows := usrdb.db.QueryRow(querys.UserUpdateQuey,
		User.FirstName,
		User.LastName,
		User.Email,
		User.ImageUrl,
		User.UserType,
		time.Now(),
		User.ID,
	)
	err := rows.Scan(
		&User.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return User, err
}

func (usrdb *userRepo) Delete(id int64) error {
	res, err := usrdb.db.Exec(querys.UserDeleteQuery, time.Now(), id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return nil
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
