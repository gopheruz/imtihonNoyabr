package main

import (
	"fmt"
	"log"

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
	data, err := strg.User().Create(&repo.UserRepo{
		FirstName: "nurmuhammad",
		LastName:  "noyabr",
		Email:     "nurmuhammad@gmail.com",
		Password:  "bilmaysan@admin",
		ImageUrl:  "google.com",
	})
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	fmt.Println(data)

	many, err := strg.User().GetAll(
		&repo.GetallUsersParams{
			Limit:  10,
			Page:   1,
			Search: "nurmuhammad",
			Sortby: "asc",
		},
	)
	if err != nil {
		log.Fatalf("failed to get all users: %v", err)
	}
	for _, user := range many.Users {
		fmt.Println(user)
	}
	fmt.Println(many.Count)
}
