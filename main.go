package main

import (
	"fmt"
	"go_crud/app/models"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DBName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	fmt.Println(models.DB)

	/*
		u := &models.User{}
		u.Name = "test"
		u.Email = "test@example.com"
		u.Password = "password"
		fmt.Println(u)

		u.CreateUser()
	*/

	u, _ := models.GetUser(1)
	fmt.Println(u)
}
