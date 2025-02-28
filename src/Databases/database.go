package databases

import (
	"fmt"

	models "github.com/Zyprush18/Kasir_Go/src/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := "root:@tcp(localhost:3306)/belajar_net_http?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connected!")

	// migration
	if err := DB.AutoMigrate(models.User{}); err != nil {
		panic(
			"Failed to migrate database!",
		)
	}
	fmt.Println("Database migrated!")

}
