package postgres_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/nurmuhammaddeveloper/imtihonNoyabr/storage/repo"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T) *repo.UserRepo {
	User, err := strg.User().Create(&repo.UserRepo{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Password:  faker.Password(),
		ImageUrl:  faker.URL(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, User)
	return User
}
func deleteUser(id int64, t *testing.T) {
	err := strg.User().Delete(id)
	require.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	n := createUser(t)
	note, err := strg.User().Get(n.ID)
	require.NoError(t, err)
	require.NotEmpty(t, note)
	deleteUser(int64(note.ID), t)
}
func TestCreateUser(t *testing.T) {
	createUser(t)
}
func TestUpdateUser(t *testing.T) {
	n := createUser(t)
	n.FirstName = faker.FirstName()
	n.LastName = faker.LastName()
	n.Email = faker.Email()
	User, err := strg.User().Update(n)
	if err != nil{
		panic(err)
	}
	fmt.Println(User)
	_, err = strg.User().Create(&repo.UserRepo{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		UserType:  "superadmin",
		Password:  faker.Password(),
		ImageUrl:  faker.URL(),
	})
	if err != nil {
		log.Fatalf("filed to create user: %v", err)
	}
	data, err := strg.User().Get(3)
	if err != nil {
		log.Fatalf("filed to get user %v", err)
	}
	fmt.Println(data)
	data, err = strg.User().Create(&repo.UserRepo{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		UserType:  "superadmin",
		Password:  faker.Password(),
		ImageUrl:  faker.URL(),
	})
	if err != nil {
		log.Fatalf("filed to create user: %v", err)
	}
	fmt.Println(data)

	data, err = strg.User().Get(3)
	if err != nil {
		log.Fatalf("filed to get user %v", err)
	}
	fmt.Println(data)
	require.NoError(t, err)
	require.NotEmpty(t, User)
	deleteUser(User.ID, t)
}
func TestDeleteUser(t *testing.T) {
	u := createUser(t)
	deleteUser(u.ID, t)
}
func TestGetAllUser(t *testing.T) {
	u := createUser(t)
	n, _ := faker.RandomInt(100)
	_, err := strg.User().GetAll(&repo.GetallUsersParams{
		Page:  int64(n[0]),
		Limit: int64(n[0]),
	})
	require.NoError(t, err)
	deleteUser(u.ID, t)
}
