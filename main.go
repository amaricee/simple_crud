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

	// UPDATE
updated := models.User{
	Name:  "Siti Aminah Updated",
	Email: "siti.updated@mail.com",
}

	ok, err := controllers.UpdateUser(2, updated)
	if err != nil {
		panic(err)
	}
	if !ok {
		fmt.Println("User not found for update")
	} else {
		fmt.Println("User updated")
	}

// DELETE
	ok, err = controllers.DeleteUser(1)
	if err != nil {
		panic(err)
	}
	if !ok {
		fmt.Println("User not found for delete")
	} else {
		fmt.Println("User deleted")
	}
}
