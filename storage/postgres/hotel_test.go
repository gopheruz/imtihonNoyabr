package postgres_test

import (
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/nurmuhammaddeveloper/imtihonNoyabr/storage/repo"
	"github.com/stretchr/testify/require"
)

func createHotel(t *testing.T) *repo.HotelRepo {
	hotel, err := strg.Hotel().Create(&repo.HotelRepo{
		HotelName:   faker.Name(),
		HotelAdress: faker.Sentence(),
		Description: faker.Sentence(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, hotel)
	return hotel
}
func deleteHotel(id int64, t *testing.T) {
	err := strg.Hotel().Delete(id)
	require.NoError(t, err)
}

func TestCreateHotel(t *testing.T) {
	createUser(t)
}
