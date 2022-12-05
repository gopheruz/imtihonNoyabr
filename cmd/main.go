package main

import (
	"fmt"
	"log"

	"github.com/bxcodec/faker/v4"
	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nurmuhammaddeveloper/imtihonNoyabr/config"
	"github.com/nurmuhammaddeveloper/imtihonNoyabr/storage"
	"github.com/nurmuhammaddeveloper/imtihonNoyabr/storage/repo"
)

func main() {
	cfg := config.Load("./")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostConfig.Host,
		cfg.PostConfig.Port,
		cfg.PostConfig.User,
		cfg.PostConfig.Password,
		cfg.PostConfig.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisConfig.RedisHost + ":" + cfg.RedisConfig.RedisPort,
	})
	fmt.Println(rdb)
	strg := storage.NewStoragePg(psqlConn)
	fmt.Println(strg)
	data, err := strg.User().Update(&repo.UserRepo{
		ID: 3,
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		Email: faker.Email(),
		ImageUrl: faker.URL(),
		UserType: "user",
	})
	if err != nil{
		log.Fatalf("filed to update user %v", err)
	}
	fmt.Println(data)
}