package main

import (
	"bwatest/users/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwatest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	userRepository := user.NewRepository(db)
	user := user.User{
		Name:       "gogi",
		Occupation: "sysadmin",
		Email:      "gogi@gmail.com",
	}
	userRepository.Save(user)
}
