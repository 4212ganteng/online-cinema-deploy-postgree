package databases

import (
	"fmt"
	"online-cinema/models"
	"online-cinema/pkg/mysql"
)

func RunMigrate() {
	if err := mysql.DB.AutoMigrate(&models.User{}, &models.Categorie{}, &models.Film{}, &models.Transaction{}); err != nil {
		fmt.Println("err")
		panic("Migration Failed")
	}
	fmt.Println("Migration Success")
}
