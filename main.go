package main

import (
	"fmt"

	"github.com/amaricee/simple_crud/config"
	"github.com/amaricee/simple_crud/controllers"
	"github.com/amaricee/simple_crud/models"
)

func main() {
	config.ConnectDB()

	// CREATE
	user := models.User{
		Name:  "Siti Aminah",
		Email: "siti@mail.com",
	}

	err := controllers.CreateUser(user)
	if err != nil {
		panic(err)
	}

	fmt.Println("User created")

	// READ
	users, err := controllers.GetAllUsers()
	if err != nil {
		panic(err)
	}

	fmt.Println("All users:")
	for _, u := range users {
		fmt.Println(u.ID, u.Name, u.Email)
	}
}
