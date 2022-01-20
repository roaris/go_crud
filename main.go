package main

import (
	"fmt"
	"go_crud/app/controllers"
	"go_crud/app/models"
)

func main() {
	fmt.Println(models.DB)

	controllers.StartMainServer()
}
