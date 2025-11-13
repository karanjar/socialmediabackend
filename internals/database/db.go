package database

import (
	"fmt"
	friendship "socialmediabackend/models/friendship"
	"socialmediabackend/models/posts"
	"socialmediabackend/models/users"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "user=dustin password=12345 dbname=socialmedia sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("unable to open  database: %v", err)
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("unable to get sql database from gorm,err: %v", err)
		panic(err)

	}

	if err := sqlDB.Ping(); err != nil {
		fmt.Printf("unable to connect database,error: %v", err)
		panic(err)
	}
	DB = db
	DB.AutoMigrate(&users.Users{}, &friendship.Friendship{}, &posts.Post{})

	fmt.Println("Database connected successfully")
}
